package azure

type Subscription struct {
	ID                   string            `json:"id"`
	AuthorizationSource  string            `json:"authorizationSource"`
	ManagedByTenants     []interface{}     `json:"managedByTenants"`
	Tags                 map[string]string `json:"tags"`
	SubscriptionID       string            `json:"subscriptionId"`
	TenantID             string            `json:"tenantId"`
	DisplayName          string            `json:"displayName"`
	State                string            `json:"state"`
	SubscriptionPolicies struct {
		LocationPlacementID string `json:"locationPlacementId"`
		QuotaID             string `json:"quotaId"`
		SpendingLimit       string `json:"spendingLimit"`
	} `json:"subscriptionPolicies"`
}

type ListSubscriptionsResponse struct {
	Value []Subscription `json:"value"`
	Count struct {
		Type  string `json:"type"`
		Value int    `json:"value"`
	} `json:"count"`
}

type GetResourceResponse struct {
	Resources []Resource `json:"value"`
}

type Resource struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Sku      map[string]string `json:"sku,omitempty"`
	Location string            `json:"location"`
	Tags     map[string]string `json:"tags,omitempty"`
	Kind     string            `json:"kind,omitempty"`
}
