package catalog

import (
	"reflect"
	"strings"
)

// WellKnownAnnotations contains a number of well-known annotations with defined semantics.
// They can be attached to catalog entities and consumed by plugins as needed.
//
// See: https://backstage.io/docs/features/software-catalog/well-known-annotations
type WellKnownAnnotations struct {
	// BackstageManagedByLocation is a so-called location reference string.
	//
	// It points to the source from which the entity was originally fetched.
	//
	// This annotation is added automatically by the catalog as it fetches the
	// data from a registered location, and is not meant to normally be written
	// by humans.
	//
	// The annotation may point to any type of generic location that the catalog
	// supports, so it cannot be relied on to always be specifically of type url,
	// nor that it even represents a single file. Note also that a single
	// location can be the source of many entities, so it represents a
	// many-to-one relationship.
	//
	//
	// The format of the value is <type>:<target>.
	//
	// Note that the target may also contain colons, so it is not advisable to
	// naively split the value on : and expecting a two-item array out of it. The
	// format of the target part is type-dependent and could conceivably even be
	// an empty string, but the separator colon is always present.
	BackstageManagedByLocation string `json:"backstage.io/managed-by-location,omitempty"`

	// BackstageManagedByOriginLocation is a location reference string (see above).
	//
	// It points to the location, whose registration lead to the creation of the
	// entity.
	//
	// In most cases, the backstage.io/managed-by-location and
	// backstage.io/managed-by-origin-location will be equal. They will be
	// different if the original location delegates to another location.
	//
	// A common case is, that a location is registered as bootstrap:bootstrap
	// which means that it is part of the app-config.yaml of a Backstage
	// installation.
	BackstageManagedByOriginLocation string `json:"backstage.io/managed-by-origin-location,omitempty"`

	// BackstageOrphan is either absent, or present with the exact string value "true".
	//
	// It should never be added manually. Instead, the catalog itself injects the
	// annotation as part of its processing loops, on entities that are found to
	// have no registered locations or config locations that keep them "active" /
	// "alive".
	//
	// For example, suppose that the user first registers a location URL pointing
	// to a Location kind entity, which in turn refers to two Component kind
	// entities in two other files nearby. The end result is that the catalog
	// contains those three entities. Now suppose that the user edits the original
	// Location entity to only refer to the first of the Component kind entities.
	// This will intentionally not lead to the other Component entity to be removed
	// from the catalog (for safety reasons). Instead, it gains this orphan marker
	// annotation, to make it clear that user action is required to completely
	// remove it, if desired.
	BackstageOrphan string `json:"backstage.io/orphan,omitempty"`

	// BackstageTechDocsRef informs where TechDocs source content is stored so
	// that it can be read and docs can be generated from it.
	//
	// Most commonly, it's written as a path, relative to the location of the
	// catalog-info.yaml itself, where the associated mkdocs.yml file can be
	// found.
	//
	// In unusual situations where the documentation for a catalog entity does
	// not live alongside the entity's source code, the value of this annotation
	// can point to an absolute URL, matching the location reference string
	// format outlined above, for example:
	// url:https://github.com/backstage/backstage/tree/master
	BackstageTechDocsRef string `json:"backstage.io/techdocs-ref,omitempty"`

	// BackstageViewURL allows customizing links from the catalog pages.
	//
	// The view URL should point to the canonical metadata YAML that governs this entity.
	BackstageViewURL string `json:"backstage.io/view-url,omitempty"`

	// BackstageEditURL allows customizing links from the catalog pages.
	//
	// The edit URL should point to the source file for the metadata.
	BackstageEditURL string `json:"backstage.io/edit-url,omitempty"`

	// BackstageSourceLocation is a Location reference that points to the source
	// code of the entity (typically a Component).
	//
	// Useful when catalog files do not get ingested from the source code
	// repository itself.
	//
	// If the URL points to a folder, it is important that it is suffixed with a
	// '/' in order for relative path resolution to work consistently.
	BackstageSourceLocation string `json:"backstage.io/source-location,omitempty"`

	// GitHubProjectSlug is the so-called slug that identifies a repository on
	// GitHub that is related to this entity.
	//
	// It is on the format <organization or owner>/<repository>, and is the same
	// as can be seen in the URL location bar of the browser when viewing that
	// repository.
	//
	// Specifying this annotation will enable GitHub related features in
	// Backstage for that entity.
	GitHubProjectSlug string `json:"github.com/project-slug,omitempty"`

	// GitHubTeamSlug is the so-called slug that identifies a team on GitHub that is
	// related to this entity.
	//
	// It is on the format <organization>/<team>, and is the same as can be seen
	// in the URL location bar of the browser when viewing that team.
	//
	// This annotation can be used on a Group entity to note that it originated
	// from that team on GitHub.
	GitHubTeamSlug string `json:"github.com/team-slug,omitempty"`

	// GitHubUserLogin is the so-called login that identifies a user on GitHub
	// that is related to this entity.
	//
	// It is on the format <username>, and is the same as can be seen in the URL
	// location bar of the browser when viewing that user.
	//
	// This annotation can be used on a User entity to note that it originated
	// from that user on GitHub.
	GitHubUserLogin string `json:"github.com/user-login,omitempty"`
}

// UnmarshalAnnotations sets the wll-known annotations from a map of annotations.
func (w *WellKnownAnnotations) UnmarshalAnnotations(annotations map[string]string) {
	wt := reflect.TypeOf(w).Elem()
	wv := reflect.ValueOf(w).Elem()
	for i := 0; i < wt.NumField(); i++ {
		field := wt.Field(i)
		if field.Type.Kind() != reflect.String {
			continue
		}
		if fieldName, _, _ := strings.Cut(field.Tag.Get("json"), ","); fieldName != "" {
			if annotation, ok := annotations[fieldName]; ok {
				wv.Field(i).SetString(annotation)
			}
		}
	}
}
