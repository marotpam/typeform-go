# typeform-go
Non-official Go library for the Typeform API https://developer.typeform.com/

## Installation
```shell
go get github.com/marotpam/typeform-go
```

## Basic usage

```go
// create a typeform client with a personal access token (see https://developer.typeform.com/get-started/personal-access-token/)
tfClient := typeform.NewDefaultClient("personalAccessToken")
svc := typeform.NewFormService(tfClient)

// create a new form
f, err := svc.Create(typeform.Form{
  Title: "a form created with typeform-go",
  Fields: []*typeform.Field{
    {
      Title: "do you like go?",
      Type:  typeform.FieldTypeYesNo,
    },
  },
})
if err != nil {
  log.Fatal(err)
}
log.Printf("typeform with id %s created\n", f.ID)

// retrieve an existing form
f, err = svc.Retrieve(f.ID)
if err != nil {
  log.Fatal(err)
}
log.Printf("typeform with id '%s' retrieved successfully\n", f.ID)

// update an existing form
f.Title = "updated form title"
f, err = svc.Update(f)
if err != nil {
  log.Fatal(err)
}
log.Printf("typeform's title is now '%s'\n", f.Title)

// list available forms filtered by params
l, err := svc.List(typeform.FormListParams{
  Search:      "form title",
  Page:        1,
  PageSize:    1,
  WorkspaceID: "workspaceID",
})
if err != nil {
  log.Fatal(err)
}
log.Printf("total number of forms: %d\n", l.TotalItems)

// delete an existing form
err = svc.Delete(f.ID)
if err != nil {
  log.Fatal(err)
}
log.Printf("typeform with id %s successfully deleted\n", f.ID)
```
