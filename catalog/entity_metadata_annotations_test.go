package catalog

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestWellKnownAnnotations_UnmarshalAnnotations(t *testing.T) {
	annotations := map[string]string{
		"backstage.io/managed-by-location":        "url:http://github.com/backstage/backstage/blob/master/catalog-info.yaml",
		"backstage.io/managed-by-origin-location": "url:http://github.com/backstage/backstage/blob/master/catalog-info.yaml",
		"backstage.io/orphan":                     "true",
		"backstage.io/techdocs-ref":               "dir:.",
		"backstage.io/view-url":                   "https://some.website/catalog-info.yaml",
		"backstage.io/edit-url":                   "https://github.com/my-org/catalog/edit/master/my-service.jsonnet",
		"backstage.io/source-location":            "url:https://github.com/my-org/my-service/",
		"github.com/project-slug":                 "backstage/backstage",
		"github.com/team-slug":                    "backstage/maintainers",
		"github.com/user-login":                   "freben",
		"example.com/unknown-annotation":          "foo",
	}
	expected := WellKnownAnnotations{
		BackstageManagedByLocation:       "url:http://github.com/backstage/backstage/blob/master/catalog-info.yaml",
		BackstageManagedByOriginLocation: "url:http://github.com/backstage/backstage/blob/master/catalog-info.yaml",
		BackstageOrphan:                  "true",
		BackstageTechDocsRef:             "dir:.",
		BackstageViewURL:                 "https://some.website/catalog-info.yaml",
		BackstageEditURL:                 "https://github.com/my-org/catalog/edit/master/my-service.jsonnet",
		BackstageSourceLocation:          "url:https://github.com/my-org/my-service/",
		GitHubProjectSlug:                "backstage/backstage",
		GitHubTeamSlug:                   "backstage/maintainers",
		GitHubUserLogin:                  "freben",
	}
	var actual WellKnownAnnotations
	actual.UnmarshalAnnotations(annotations)
	assert.DeepEqual(t, expected, actual)
}
