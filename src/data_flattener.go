package main

func FlatenPayloadSearchParam(payloadData payload) payloadSearchParam {
	var result payloadSearchParam
	result.Company = payloadData.Company
	result.Description = payloadData.Description
	result.License = payloadData.License
	result.Source = payloadData.Source
	result.Title = payloadData.Title
	result.Version = payloadData.Version
	result.Website = payloadData.Website

	for _, v := range payloadData.Maintainers {
		result.MaintainersNames = append(result.MaintainersNames, v.Name)
		result.MaintainersEmails = append(result.MaintainersEmails, v.Email)
	}

	return result
}
