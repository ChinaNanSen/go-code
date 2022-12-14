//
// Copyright 2020 Joyent, Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//

package compute_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"testing"

	"github.com/abdullin/seq"
	triton "github.com/joyent/triton-go"
	"github.com/joyent/triton-go/compute"
	"github.com/joyent/triton-go/testutils"
	"github.com/pkg/errors"
)

var (
	fakeImageID = "2b683a82-a066-11e3-97ab-2faa44701c5a"
	localImage  *compute.Image
)

func TestAccImagesList(t *testing.T) {
	const stateKey = "images"

	testutils.AccTest(t, testutils.TestCase{
		Steps: []testutils.Step{

			&testutils.StepClient{
				StateBagKey: stateKey,
				CallFunc: func(config *triton.ClientConfig) (interface{}, error) {
					return compute.NewClient(config)
				},
			},

			&testutils.StepAPICall{
				StateBagKey: stateKey,
				CallFunc: func(client interface{}) (interface{}, error) {
					c := client.(*compute.ComputeClient)
					ctx := context.Background()
					input := &compute.ListImagesInput{}
					return c.Images().List(ctx, input)
				},
			},

			&testutils.StepAssertFunc{
				AssertFunc: func(state testutils.TritonStateBag) error {
					images, ok := state.GetOk(stateKey)
					if !ok {
						return fmt.Errorf("State key %q not found", stateKey)
					}

					if len(images.([]*compute.Image)) == 0 {
						return fmt.Errorf("No images returned by list images")
					}

					localImage = images.([]*compute.Image)[0]

					if localImage.Name == "" {
						return fmt.Errorf("Returned image should have a name")
					}

					if localImage.Owner == "" {
						return fmt.Errorf("Returned image should have an owner")
					}

					return nil
				},
			},
		},
	})
}

func TestAccImagesListInput(t *testing.T) {
	const stateKey = "images"
	const imageKey = "image"

	testutils.AccTest(t, testutils.TestCase{
		Steps: []testutils.Step{

			&testutils.StepClient{
				StateBagKey: stateKey,
				CallFunc: func(config *triton.ClientConfig) (interface{}, error) {
					return compute.NewClient(config)
				},
			},

			&testutils.StepAPICall{
				StateBagKey: stateKey,
				CallFunc: func(client interface{}) (interface{}, error) {
					c := client.(*compute.ComputeClient)
					ctx := context.Background()
					input := &compute.ListImagesInput{
						Name:    localImage.Name,
						Type:    localImage.Type,
						Version: localImage.Version,
					}
					return c.Images().List(ctx, input)
				},
			},

			&testutils.StepAssertFunc{
				AssertFunc: func(state testutils.TritonStateBag) error {
					images, ok := state.GetOk(stateKey)
					if !ok {
						return fmt.Errorf("State key %q not found", stateKey)
					}

					for _, image := range images.([]*compute.Image) {
						if image.ID == localImage.ID {
							state.Put(imageKey, image)
							return nil
						}
					}

					return fmt.Errorf("Did not find image %q", localImage.ID)
				},
			},

			&testutils.StepAssert{
				StateBagKey: imageKey,
				Assertions: seq.Map{
					"id":      localImage.ID,
					"name":    localImage.Name,
					"owner":   localImage.Owner,
					"type":    localImage.Type,
					"version": localImage.Version,
				},
			},
		},
	})
}

func TestAccImagesGet(t *testing.T) {
	const stateKey = "image"

	testutils.AccTest(t, testutils.TestCase{
		Steps: []testutils.Step{

			&testutils.StepClient{
				StateBagKey: stateKey,
				CallFunc: func(config *triton.ClientConfig) (interface{}, error) {
					return compute.NewClient(config)
				},
			},

			&testutils.StepAPICall{
				StateBagKey: stateKey,
				CallFunc: func(client interface{}) (interface{}, error) {
					c := client.(*compute.ComputeClient)
					ctx := context.Background()
					input := &compute.GetImageInput{
						ImageID: localImage.ID,
					}
					return c.Images().Get(ctx, input)
				},
			},

			&testutils.StepAssert{
				StateBagKey: stateKey,
				Assertions: seq.Map{
					"id":      localImage.ID,
					"name":    localImage.Name,
					"owner":   localImage.Owner,
					"type":    localImage.Type,
					"version": localImage.Version,
				},
			},
		},
	})
}

func TestDeleteImage(t *testing.T) {
	computeClient := MockComputeClient()

	do := func(ctx context.Context, cc *compute.ComputeClient) error {
		defer testutils.DeactivateClient()

		return cc.Images().Delete(ctx, &compute.DeleteImageInput{
			ImageID: fakeImageID,
		})
	}

	t.Run("successful", func(t *testing.T) {
		testutils.RegisterResponder("DELETE", path.Join("/", accountURL, "images", fakeImageID), deleteImageSuccess)

		err := do(context.Background(), computeClient)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("error", func(t *testing.T) {
		testutils.RegisterResponder("DELETE", path.Join("/", accountURL, "images", fakeImageID), deleteImageError)

		err := do(context.Background(), computeClient)
		if err == nil {
			t.Fatal(err)
		}

		if !strings.Contains(err.Error(), "unable to delete image") {
			t.Errorf("expected error to equal testError: found %s", err)
		}
	})
}

func TestGetImage(t *testing.T) {
	computeClient := MockComputeClient()

	do := func(ctx context.Context, cc *compute.ComputeClient) (*compute.Image, error) {
		defer testutils.DeactivateClient()

		image, err := cc.Images().Get(ctx, &compute.GetImageInput{
			ImageID: fakeImageID,
		})
		if err != nil {
			return nil, err
		}
		return image, nil
	}

	t.Run("successful", func(t *testing.T) {
		testutils.RegisterResponder("GET", path.Join("/", accountURL, "images", fakeImageID), getImageSuccess)

		resp, err := do(context.Background(), computeClient)
		if err != nil {
			t.Fatal(err)
		}

		if resp == nil {
			t.Fatalf("Expected an output but got nil")
		}
	})

	t.Run("eof", func(t *testing.T) {
		testutils.RegisterResponder("GET", path.Join("/", accountURL, "images", fakeImageID), getImageEmpty)

		_, err := do(context.Background(), computeClient)
		if err == nil {
			t.Fatal(err)
		}

		if !strings.Contains(err.Error(), "EOF") {
			t.Errorf("expected error to contain EOF: found %s", err)
		}
	})

	t.Run("bad_decode", func(t *testing.T) {
		testutils.RegisterResponder("GET", path.Join("/", accountURL, "images", fakeImageID), getImageBadDecode)

		_, err := do(context.Background(), computeClient)
		if err == nil {
			t.Fatal(err)
		}

		if !strings.Contains(err.Error(), "invalid character") {
			t.Errorf("expected decode to fail: found %s", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		testutils.RegisterResponder("GET", path.Join("/", accountURL, "images", fakeImageID), getImageError)

		resp, err := do(context.Background(), computeClient)
		if err == nil {
			t.Fatal(err)
		}
		if resp != nil {
			t.Error("expected resp to be nil")
		}

		if !strings.Contains(err.Error(), "unable to get image") {
			t.Errorf("expected error to equal testError: found %s", err)
		}
	})
}

func TestListImages(t *testing.T) {
	computeClient := MockComputeClient()

	do := func(ctx context.Context, cc *compute.ComputeClient) ([]*compute.Image, error) {
		defer testutils.DeactivateClient()

		images, err := cc.Images().List(ctx, &compute.ListImagesInput{})
		if err != nil {
			return nil, err
		}
		return images, nil
	}

	t.Run("successful", func(t *testing.T) {
		testutils.RegisterResponder("GET", path.Join("/", accountURL, "images"), listImagesSuccess)

		resp, err := do(context.Background(), computeClient)
		if err != nil {
			t.Fatal(err)
		}

		if resp == nil {
			t.Fatalf("Expected an output but got nil")
		}
	})

	t.Run("eof", func(t *testing.T) {
		testutils.RegisterResponder("GET", path.Join("/", accountURL, "images"), listImagesEmpty)

		_, err := do(context.Background(), computeClient)
		if err == nil {
			t.Fatal(err)
		}

		if !strings.Contains(err.Error(), "EOF") {
			t.Errorf("expected error to contain EOF: found %s", err)
		}
	})

	t.Run("bad_decode", func(t *testing.T) {
		testutils.RegisterResponder("GET", path.Join("/", accountURL, "images"), listImagesBadDecode)

		_, err := do(context.Background(), computeClient)
		if err == nil {
			t.Fatal(err)
		}

		if !strings.Contains(err.Error(), "invalid character") {
			t.Errorf("expected decode to fail: found %s", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		testutils.RegisterResponder("GET", path.Join("/", accountURL, "images"), listImagesError)

		resp, err := do(context.Background(), computeClient)
		if err == nil {
			t.Fatal(err)
		}
		if resp != nil {
			t.Error("expected resp to be nil")
		}

		if !strings.Contains(err.Error(), "unable to list images") {
			t.Errorf("expected error to equal testError: found %v", err)
		}
	})
}

func TestCreateImageFromMachine(t *testing.T) {
	computeClient := MockComputeClient()

	do := func(ctx context.Context, cc *compute.ComputeClient) (*compute.Image, error) {
		defer testutils.DeactivateClient()

		image, err := cc.Images().CreateFromMachine(ctx, &compute.CreateImageFromMachineInput{
			MachineID: "a44f2b9b-e7af-f548-b0ba-4d9270423f1a",
			Name:      "my-custom-image",
			Version:   "1.0.0",
		})
		if err != nil {
			return nil, err
		}
		return image, nil
	}

	t.Run("successful", func(t *testing.T) {
		testutils.RegisterResponder("POST", path.Join("/", accountURL, "images"), createImageFromMachineSuccess)

		_, err := do(context.Background(), computeClient)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("error", func(t *testing.T) {
		testutils.RegisterResponder("POST", path.Join("/", accountURL, "images"), createImageFromMachineError)

		_, err := do(context.Background(), computeClient)
		if err == nil {
			t.Fatal(err)
		}

		if !strings.Contains(err.Error(), "unable to create image from machine") {
			t.Errorf("expected error to equal testError: found %s", err)
		}
	})
}

func TestUpdateImage(t *testing.T) {
	computeClient := MockComputeClient()

	do := func(ctx context.Context, cc *compute.ComputeClient) (*compute.Image, error) {
		defer testutils.DeactivateClient()

		image, err := cc.Images().Update(ctx, &compute.UpdateImageInput{
			ImageID: fakeImageID,
			Version: "1.0.1",
		})
		if err != nil {
			return nil, err
		}
		return image, nil
	}

	t.Run("successful", func(t *testing.T) {

		testutils.RegisterResponder("POST", fmt.Sprintf("/%s/images/%s?action=update", accountURL, fakeImageID), updateImageSuccess)

		_, err := do(context.Background(), computeClient)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("error", func(t *testing.T) {
		testutils.RegisterResponder("POST", fmt.Sprintf("/%s/images/not-a-real-image?action=update", accountURL), updateImageError)

		_, err := do(context.Background(), computeClient)
		if err == nil {
			t.Fatal(err)
		}

		if !strings.Contains(err.Error(), "unable to update image") {
			t.Errorf("expected error to equal testError: found %s", err)
		}
	})
}

func TestExportImage(t *testing.T) {
	computeClient := MockComputeClient()

	do := func(ctx context.Context, cc *compute.ComputeClient) (*compute.MantaLocation, error) {
		defer testutils.DeactivateClient()

		location, err := cc.Images().Export(ctx, &compute.ExportImageInput{
			ImageID:   fakeImageID,
			MantaPath: "/stor/images/myimages",
		})
		if err != nil {
			return nil, err
		}
		return location, nil
	}

	t.Run("successful", func(t *testing.T) {

		testutils.RegisterResponder("POST", fmt.Sprintf("/%s/images/%s?action=export", accountURL, fakeImageID), exportImageSuccess)

		_, err := do(context.Background(), computeClient)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("error", func(t *testing.T) {
		testutils.RegisterResponder("POST", fmt.Sprintf("/%s/images/%s?action=export", accountURL, fakeImageID), exportImageError)

		_, err := do(context.Background(), computeClient)
		if err == nil {
			t.Fatal(err)
		}

		if !strings.Contains(err.Error(), "unable to export image") {
			t.Errorf("expected error to equal testError: found %s", err)
		}
	})
}

func deleteImageSuccess(req *http.Request) (*http.Response, error) {
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	return &http.Response{
		StatusCode: http.StatusNoContent,
		Header:     header,
	}, nil
}

func deleteImageError(req *http.Request) (*http.Response, error) {
	return nil, errors.New("unable to delete image")
}

func getImageSuccess(req *http.Request) (*http.Response, error) {
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	body := strings.NewReader(`{
  "id": "2b683a82-a066-11e3-97ab-2faa44701c5a",
  "name": "base",
  "version": "13.4.0",
  "os": "smartos",
  "requirements": {},
  "type": "zone-dataset",
  "description": "A 32-bit SmartOS image with just essential packages installed. Ideal for users who are comfortable with setting up their own environment and tools.",
  "files": [
	{
	  "compression": "gzip",
	  "sha1": "3bebb6ae2cdb26eef20cfb30fdc4a00a059a0b7b",
	  "size": 110742036
	}
  ],
  "tags": {
	"role": "os",
	"group": "base-32"
  },
  "homepage": "https://docs.joyent.com/images/smartos/base",
  "published_at": "2014-02-28T10:50:42Z",
  "owner": "930896af-bf8c-48d4-885c-6573a94b1853",
  "public": true,
  "state": "active"
}
`)

	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     header,
		Body:       ioutil.NopCloser(body),
	}, nil
}

func getImageBadDecode(req *http.Request) (*http.Response, error) {
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	body := strings.NewReader(`{
  "id": "2b683a82-a066-11e3-97ab-2faa44701c5a",
  "name": "base",
  "version": "13.4.0",
  "os": "smartos",
  "requirements": {},
  "type": "zone-dataset",
  "description": "A 32-bit SmartOS image with just essential packages installed. Ideal for users who are comfortable with setting up their own environment and tools.",
  "files": [
	{
	  "compression": "gzip",
	  "sha1": "3bebb6ae2cdb26eef20cfb30fdc4a00a059a0b7b",
	  "size": 110742036
	}
  ],
  "tags": {
	"role": "os",
	"group": "base-32"
  },
  "homepage": "https://docs.joyent.com/images/smartos/base",
  "published_at": "2014-02-28T10:50:42Z",
  "owner": "930896af-bf8c-48d4-885c-6573a94b1853",
  "public": true,
  "state": "active",
}`)

	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     header,
		Body:       ioutil.NopCloser(body),
	}, nil
}

func getImageEmpty(req *http.Request) (*http.Response, error) {
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     header,
		Body:       ioutil.NopCloser(strings.NewReader("")),
	}, nil
}

func getImageError(req *http.Request) (*http.Response, error) {
	return nil, errors.New("unable to get image")
}

func listImagesSuccess(req *http.Request) (*http.Response, error) {
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	body := strings.NewReader(`[
{
	"id": "2b683a82-a066-11e3-97ab-2faa44701c5a",
	"name": "base",
	"version": "13.4.0",
	"os": "smartos",
	"requirements": {},
	"type": "zone-dataset",
	"description": "A 32-bit SmartOS image with just essential packages installed. Ideal for users who are comfortable with setting up their own environment and tools.",
	"files": [
	  {
		"compression": "gzip",
		"sha1": "3bebb6ae2cdb26eef20cfb30fdc4a00a059a0b7b",
		"size": 110742036
	  }
	],
	"tags": {
	  "role": "os",
	  "group": "base-32"
	},
	"homepage": "https://docs.joyent.com/images/smartos/base",
	"published_at": "2014-02-28T10:50:42Z",
	"owner": "930896af-bf8c-48d4-885c-6573a94b1853",
	"public": true,
	"state": "active"
  }
]`)

	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     header,
		Body:       ioutil.NopCloser(body),
	}, nil
}

func listImagesEmpty(req *http.Request) (*http.Response, error) {
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     header,
		Body:       ioutil.NopCloser(strings.NewReader("")),
	}, nil
}

func listImagesBadDecode(req *http.Request) (*http.Response, error) {
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	body := strings.NewReader(`[{
	"id": "2b683a82-a066-11e3-97ab-2faa44701c5a",
	"name": "base",
	"version": "13.4.0",
	"os": "smartos",
	"requirements": {},
	"type": "zone-dataset",
	"description": "A 32-bit SmartOS image with just essential packages installed. Ideal for users who are comfortable with setting up their own environment and tools.",
	"files": [
	  {
		"compression": "gzip",
		"sha1": "3bebb6ae2cdb26eef20cfb30fdc4a00a059a0b7b",
		"size": 110742036
	  }
	],
	"tags": {
	  "role": "os",
	  "group": "base-32"
	},
	"homepage": "https://docs.joyent.com/images/smartos/base",
	"published_at": "2014-02-28T10:50:42Z",
	"owner": "930896af-bf8c-48d4-885c-6573a94b1853",
	"public": true,
	"state": "active",
  }]`)

	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     header,
		Body:       ioutil.NopCloser(body),
	}, nil
}

func listImagesError(req *http.Request) (*http.Response, error) {
	return nil, errors.New("unable to list images")
}

func createImageFromMachineSuccess(req *http.Request) (*http.Response, error) {
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	body := strings.NewReader(`{
	"id": "62306cd7-7b8a-c5dd-d44e-8491c83b9974",
	"name": "my-custom-image",
	"version": "1.2.3",
	"requirements": {},
	"owner": "47034e57-42d1-0342-b302-00db733e8c8a",
	"public": false,
	"state": "active"
}
`)

	return &http.Response{
		StatusCode: http.StatusCreated,
		Header:     header,
		Body:       ioutil.NopCloser(body),
	}, nil
}

func createImageFromMachineError(req *http.Request) (*http.Response, error) {
	return nil, errors.New("unable to create image from machine")
}

func updateImageSuccess(req *http.Request) (*http.Response, error) {
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	body := strings.NewReader(`{
  "id": "2b683a82-a066-11e3-97ab-2faa44701c5a",
  "name": "my-custom-image",
  "version": "1.0.1",
  "os": "smartos",
  "requirements": {},
  "type": "zone-dataset",
  "published_at": "2013-11-25T17:44:54Z",
  "owner": "47034e57-42d1-0342-b302-00db733e8c8a",
  "public": true,
  "state": "active"
}
`)

	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     header,
		Body:       ioutil.NopCloser(body),
	}, nil
}

func updateImageError(req *http.Request) (*http.Response, error) {
	return nil, errors.New("unable to update image")
}

func exportImageSuccess(req *http.Request) (*http.Response, error) {
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	body := strings.NewReader(`{
  "manta_url": "https://us-east.manta.joyent.com",
  "image_path": "/user/stor/my-image.zfs.gz",
  "manifest_path": "/user/stor/my-image.imgmanifest"
}
`)

	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     header,
		Body:       ioutil.NopCloser(body),
	}, nil
}

func exportImageError(req *http.Request) (*http.Response, error) {
	return nil, errors.New("unable to export image")
}
