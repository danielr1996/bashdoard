package types

type DashboardEntry struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Icon     string `json:"icon"`
	Id       string `json:"id"`
	Provider string `json:"provider"`
}

func (d *DashboardEntry) Equals(entry DashboardEntry) bool {
	return d.Name == entry.Name &&
		d.Url == entry.Url &&
		d.Icon == entry.Icon &&
		d.Id == entry.Id &&
		d.Provider == entry.Provider
}
