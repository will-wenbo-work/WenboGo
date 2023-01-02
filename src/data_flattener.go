package main

func FlatenEvent2EventSearchParam(eventdata event) eventSearchParam {
	var result eventSearchParam
	result.Company = eventdata.Company
	result.Description = eventdata.Description
	result.License = eventdata.License
	result.Source = eventdata.Source
	result.Title = eventdata.Title
	result.Version = eventdata.Version
	result.Website = eventdata.Website

	for _, v := range eventdata.Maintainers {
		result.MaintainersNames = append(result.MaintainersNames, v.Name)
		result.MaintainersEmails = append(result.MaintainersEmails, v.Email)
	}

	return result
}
