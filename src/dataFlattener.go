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
	result.maintainersName = eventdata.Maintainers.Name
	result.MaintainersEmail = eventdata.Maintainers.Email

	return result
}
