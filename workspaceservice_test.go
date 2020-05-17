package typeform_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marotpam/typeform-go"
)

func TestCreateWorkspace(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewWorkspaceService(newFakeServerClient(t))

		f, err := svc.Create(typeform.Workspace{
			Name: "my new workspace",
		})
		assert.Nil(t, err)

		assert.Equal(t, workspaceIDCreatedWorkspace, f.ID)
	})
}

func TestRetrieveWorkspace(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewWorkspaceService(newFakeServerClient(t))

		f, err := svc.Retrieve(workspaceIDRetrievedWorkspace)
		assert.Nil(t, err)

		assert.Equal(t, workspaceIDRetrievedWorkspace, f.ID)
	})
	t.Run("not found error", func(t *testing.T) {
		svc := typeform.NewWorkspaceService(newFakeServerClient(t))

		_, err := svc.Retrieve("unknownWorkspaceID")
		assert.NotNil(t, err)

		tfErr, ok := err.(typeform.Error)
		assert.True(t, ok)

		assert.Equal(t, typeform.CodeWorkspaceNotFound, tfErr.Code)
	})
}

func TestListWorkspaces(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewWorkspaceService(newFakeServerClient(t))

		list, err := svc.List()
		assert.Nil(t, err)

		assert.Len(t, list.Items, 1)
	})
}

func TestDeleteWorkspace(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewWorkspaceService(newFakeServerClient(t))

		assert.Nil(t, svc.Delete(workspaceIDDeletedWorkspace))
	})
}
