---
title: Overview
description: Azure Quick Review &mdash; Analyze Azure resources and identify whether they comply with Azure's best practices and recommendations.
weight: 1
---

**Azure Quick Review (azqr)** is a powerful command-line interface (CLI) tool that specializes in analyzing Azure resources to ensure compliance with Azure's best practices and recommendations. Its main objective is to offer users a comprehensive overview of their Azure resources, allowing them to easily identify any non-compliant configurations or areas for improvement.

## Azure Quick Review Recommendations

**Azure Quick Review (azqr)** scans your resources with 2 types of recommendations:

* **Azure Resource Graph (ARG)** queries provided by the [Azure Proactive Resiliency Library v2 (APRL)](https://aka.ms/aprl) project.
* **Azure Resource Manager (ARM)** queries built  with the Golang SDK

To learn more about the recommendations used by **Azure Quick Review (azqr)**, you can refer to the documentation available [here](https://azure.github.io/azqr/docs/recommendations/).

## Scan Results

The output generated by **Azure Quick Review (azqr)** is written by default to an Excel file, which contains the following sheets:

* **Recommendations**: a list with all recommendations with the number of resources that are impacted. You can use this table as an action plan to improve the compliance of your resources.
* **ImpactedResources**: a list with all resources that are impacted. You can use this table to identify resources that have issues that need to be addressed.
* **ResourceTypes**: a list of impacted resource types.
* **Inventory**: a list of all resources scanned by the tool. Here you'll find details such as SKU, Tier, Kind or calculated SLA.
* **Advisor**: a list of recommendations provided by Azure Advisor.
* **DefenderRecommendations**: a list of recommendations provided by Microsoft Defender for Cloud.
* **OutOfScope**: a list of resources that were not scanned.
* **Defender**: a list of Microsoft Defender for Cloud plans and their tiers.
* **Costs**: a list of costs associated with the scanned subscription for the last 3 months.


> By default, Azure Quick Review (azqr) obfuscates the Subscription Ids in the output to ensure the protection of sensitive information and maintain data privacy and security. If you want to display the Subscription Ids without obfuscation, you can use the `--mask=false` flag when executing the tool.

> Azure Quick Review can also generate an csv files with the same information as the excel. To generate the csv files, you can use the `--csv` flag when running the tool.

> A Power BI template is also available to help you visualize the results generated by Azure Quick Review. You can create the template running Azure Quick Review with the `pbi` command and then loading the excel file generated by the tool.

## Supported Azure Services

**Azure Quick Review (azqr)** currently supports the following Azure services:

{{% include "./static/types.txt" %}}

## Code of Conduct

This project has adopted the [Microsoft Open Source Code of Conduct](CODE_OF_CONDUCT.md)

## Trademark Notice

> **Trademarks** This project may contain trademarks or logos for projects, products, or services. Authorized use of Microsoft trademarks or logos is subject to and must follow Microsoft’s Trademark & Brand Guidelines. Use of Microsoft trademarks or logos in modified versions of this project must not cause confusion or imply Microsoft sponsorship. Any use of third-party trademarks or logos are subject to those third-party’s policies.
