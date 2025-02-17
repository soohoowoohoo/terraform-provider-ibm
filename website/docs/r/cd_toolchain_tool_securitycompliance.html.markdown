---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_securitycompliance"
description: |-
  Manages cd_toolchain_tool_securitycompliance.
subcategory: "CD Toolchain"
---

# ibm_cd_toolchain_tool_securitycompliance

~> **Beta:** This resource is in Beta, and is subject to change.

Provides a resource for cd_toolchain_tool_securitycompliance. This allows cd_toolchain_tool_securitycompliance to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_securitycompliance" "cd_toolchain_tool_securitycompliance" {
  parameters {
		name = "name"
		evidence_repo_name = "evidence_repo_name"
		trigger_scan = "disabled"
		location = "location"
		evidence_namespace = "evidence_namespace"
		api_key = "api_key"
		scope = "scope"
		profile = "profile"
  }
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Name of tool.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.
* `parameters` - (Required, List) Unique key-value pairs representing parameters to be used to create the tool.
Nested scheme for **parameters**:
	* `api_key` - (Optional, String) The IBM Cloud API key is used to access the Security and Compliance API. You can obtain your API key with 'ibmcloud iam api-key-create' or via the console at https://cloud.ibm.com/iam#/apikeys by clicking **Create API key** (Each API key only can be viewed once).
	  * Constraints: The value must match regular expression `/\\S/`.
	* `evidence_namespace` - (Optional, String) The kind of pipeline evidence to be displayed in Security and Compliance Center for this toolchain. The evidence locker will be searched for CD (Continuous Deployment) pipeline evidence, or for CC (Continuous Compliance) pipeline evidence.
	* `evidence_repo_name` - (Required, String) To collect and store evidence for all tasks performed, a Git repository is required as an evidence locker.
	* `location` - (Optional, String)
	* `name` - (Required, String) Give this tool integration a name, for example: my-security-compliance.
	* `profile` - (Optional, String) Select an existing profile, where a profile is a collection of security controls. [Learn more.](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles) ![](https://cloud.ibm.com/media/docs/images/icons/launch-glyph.svg).
	* `scope` - (Optional, String) Select an existing scope name to narrow the focus of the validation scan. [Learn more.](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-scopes) ![](https://cloud.ibm.com/media/docs/images/icons/launch-glyph.svg).
	* `trigger_info` - (Optional, Map) trigger_info.
	* `trigger_scan` - (Optional, String) Enabling trigger validation scans provides details for a pipeline task to trigger a scan.
	  * Constraints: Allowable values are: `disabled`, `enabled`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind the tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cd_toolchain_tool_securitycompliance.
* `crn` - (String) Tool CRN.
* `href` - (String) URI representing the tool.
* `referent` - (List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (String) URI representing the this resource through an API.
	* `ui_href` - (String) URI representing the this resource through the UI.
* `resource_group_id` - (String) Resource group where tool can be found.
* `state` - (String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.
* `toolchain_crn` - (String) CRN of toolchain which the tool is bound to.
* `tool_id` - (String) Tool ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `updated_at` - (String) Latest tool update timestamp.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_cd_toolchain_tool_securitycompliance` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

```
<toolchain_id>/<tool_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind the tool to.
* `tool_id`: A string. ID of the tool bound to the toolchain.

# Syntax
```
$ terraform import ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance <toolchain_id>/<tool_id>
```
