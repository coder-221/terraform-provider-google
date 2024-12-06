// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package networkservices_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccNetworkServicesAuthzExtension_networkServicesAuthzExtensionBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesAuthzExtensionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesAuthzExtension_networkServicesAuthzExtensionBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_authz_extension.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "service", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkServicesAuthzExtension_networkServicesAuthzExtensionBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                  = "tf-test-authz-service%{random_suffix}"
  project               = "%{project}"
  region                = "us-west1"

  protocol              = "HTTP2"
  load_balancing_scheme = "INTERNAL_MANAGED"
  port_name             = "grpc"
}

resource "google_network_services_authz_extension" "default" {
  name     = "tf-test-my-authz-ext%{random_suffix}"
  project  = "%{project}"
  location = "us-west1"

  description           = "my description"
  load_balancing_scheme = "INTERNAL_MANAGED"
  authority             = "ext11.com"
  service               = google_compute_region_backend_service.default.self_link
  timeout               = "0.1s"
  fail_open             = false
  forward_headers       = ["Authorization"]
}
`, context)
}

func testAccCheckNetworkServicesAuthzExtensionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_services_authz_extension" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkServicesBasePath}}projects/{{project}}/locations/{{location}}/authzExtensions/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("NetworkServicesAuthzExtension still exists at %s", url)
			}
		}

		return nil
	}
}