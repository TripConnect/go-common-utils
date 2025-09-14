package common

import (
	"sort"
	"strings"

	"github.com/gofrs/uuid/v5"
	"github.com/tripconnect/go-common-utils/helper"
)

func BuildUUID(items ...string) uuid.UUID {
	sort.Slice(items, func(i, j int) bool {
		return items[i] > items[j]
	})

	sortedJoin := strings.Join(items, "-")
	namespaceString, _ := helper.ReadConfig[string]("uuid.namespace")
	namespaceUUID, _ := uuid.FromString(namespaceString)
	result := uuid.NewV5(namespaceUUID, sortedJoin)
	return result
}
