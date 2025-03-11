package sagemcomfast5359

import "time"

type TopologyResponse struct {
	ErrResponse
	Status []Status `json:"status"`
}

type Names struct {
	Name   string `json:"Name"`
	Source string `json:"Source"`
	Suffix string `json:"Suffix"`
	ID     string `json:"Id"`
}
type IPv4Address struct {
	Address       string `json:"Address"`
	Status        string `json:"Status"`
	Scope         string `json:"Scope"`
	AddressSource string `json:"AddressSource"`
	Reserved      bool   `json:"Reserved"`
	ID            string `json:"Id"`
}
type IPv6Address struct {
	Address       string `json:"Address"`
	Status        string `json:"Status"`
	Scope         string `json:"Scope"`
	AddressSource string `json:"AddressSource"`
	ID            string `json:"Id"`
}
type Wps struct {
	Enable            bool `json:"Enable"`
	PairingInProgress bool `json:"PairingInProgress"`
}
type DeviceTypes struct {
	Type   string `json:"Type"`
	Source string `json:"Source"`
	ID     string `json:"Id"`
}
type Bdd struct {
	CloudVersion         string    `json:"CloudVersion"`
	BDDRequestsSent      int       `json:"BDDRequestsSent"`
	BDDRequestsAnswered  int       `json:"BDDRequestsAnswered"`
	BDDRequestsFailed    int       `json:"BDDRequestsFailed"`
	BDDLastSyncTimestamp time.Time `json:"BDDLastSyncTimestamp"`
	DeviceName           string    `json:"DeviceName"`
	DeviceType           string    `json:"DeviceType"`
	ModelName            string    `json:"ModelName"`
	OperatingSystem      string    `json:"OperatingSystem"`
	SoftwareVersion      string    `json:"SoftwareVersion"`
	Manufacturer         string    `json:"Manufacturer"`
	MACVendor            string    `json:"MACVendor"`
	DeviceCategory       string    `json:"DeviceCategory"`
	PerformManualLookup  bool      `json:"PerformManualLookup"`
}
type Priority struct {
	Configuration string `json:"Configuration"`
	Type          string `json:"Type"`
}
type Security struct {
	Algorithm string `json:"Algorithm"`
	Score     int    `json:"Score"`
	Scores    []any  `json:"Scores"`
}
type SSWSta struct {
	SupportedStandards           string `json:"SupportedStandards"`
	Supports24GHz                bool   `json:"Supports24GHz"`
	Supports5GHz                 bool   `json:"Supports5GHz"`
	Supports6GHz                 bool   `json:"Supports6GHz"`
	ReconnectClass               string `json:"ReconnectClass"`
	FailedSteerCount             int    `json:"FailedSteerCount"`
	SuccessSteerCount            int    `json:"SuccessSteerCount"`
	AvgSteeringTime              int    `json:"AvgSteeringTime"`
	SupportedUNIIBands           string `json:"SupportedUNIIBands"`
	VendorSpecificElementOUIList string `json:"VendorSpecificElementOUIList"`
}
type WANAccess struct {
	BlockedReasons string `json:"BlockedReasons"`
}
type MDNSService struct {
	Name        string `json:"Name"`
	ServiceName string `json:"ServiceName"`
	Domain      string `json:"Domain"`
	Port        string `json:"Port"`
	Text        string `json:"Text"`
	ID          string `json:"Id"`
}
type Device struct {
	Key                          string        `json:"Key"`
	DiscoverySource              string        `json:"DiscoverySource"`
	Name                         string        `json:"Name"`
	DeviceType                   string        `json:"DeviceType"`
	Active                       bool          `json:"Active"`
	Tags                         string        `json:"Tags"`
	FirstSeen                    time.Time     `json:"FirstSeen"`
	LastConnection               time.Time     `json:"LastConnection"`
	LastChanged                  time.Time     `json:"LastChanged"`
	Master                       string        `json:"Master"`
	VendorClassID                string        `json:"VendorClassID"`
	UserClassID                  string        `json:"UserClassID"`
	ClientID                     string        `json:"ClientID"`
	SerialNumber                 string        `json:"SerialNumber"`
	ProductClass                 string        `json:"ProductClass"`
	Oui                          string        `json:"OUI"`
	DHCPOption55                 string        `json:"DHCPOption55"`
	ModelName                    string        `json:"ModelName,omitempty"`
	OperatingSystem              string        `json:"OperatingSystem,omitempty"`
	SoftwareVersion              string        `json:"SoftwareVersion,omitempty"`
	Manufacturer                 string        `json:"Manufacturer,omitempty"`
	DeviceCategory               string        `json:"DeviceCategory,omitempty"`
	CatalogID                    string        `json:"CatalogID,omitempty"`
	IPAddress                    string        `json:"IPAddress"`
	IPAddressSource              string        `json:"IPAddressSource"`
	Location                     string        `json:"Location"`
	PhysAddress                  string        `json:"PhysAddress"`
	Layer2Interface              string        `json:"Layer2Interface"`
	InterfaceName                string        `json:"InterfaceName"`
	MACVendor                    string        `json:"MACVendor"`
	Owner                        string        `json:"Owner"`
	UniqueID                     string        `json:"UniqueID"`
	SignalStrength               int           `json:"SignalStrength"`
	SignalNoiseRatio             int           `json:"SignalNoiseRatio"`
	LastDataDownlinkRate         int           `json:"LastDataDownlinkRate"`
	LastDataUplinkRate           int           `json:"LastDataUplinkRate"`
	EncryptionMode               string        `json:"EncryptionMode"`
	LinkBandwidth                string        `json:"LinkBandwidth"`
	SecurityModeEnabled          string        `json:"SecurityModeEnabled"`
	HtCapabilities               string        `json:"HtCapabilities"`
	VhtCapabilities              string        `json:"VhtCapabilities"`
	HeCapabilities               string        `json:"HeCapabilities"`
	SupportedMCS                 string        `json:"SupportedMCS"`
	AuthenticationState          bool          `json:"AuthenticationState"`
	OperatingStandard            string        `json:"OperatingStandard"`
	OperatingFrequencyBand       string        `json:"OperatingFrequencyBand"`
	AvgSignalStrengthByChain     int           `json:"AvgSignalStrengthByChain"`
	MaxBandwidthSupported        string        `json:"MaxBandwidthSupported"`
	MaxDownlinkRateSupported     int           `json:"MaxDownlinkRateSupported"`
	MaxDownlinkRateReached       int           `json:"MaxDownlinkRateReached"`
	DownlinkMCS                  int           `json:"DownlinkMCS"`
	DownlinkBandwidth            int           `json:"DownlinkBandwidth"`
	DownlinkShortGuard           bool          `json:"DownlinkShortGuard"`
	UplinkMCS                    int           `json:"UplinkMCS"`
	UplinkBandwidth              int           `json:"UplinkBandwidth"`
	UplinkShortGuard             bool          `json:"UplinkShortGuard"`
	MaxUplinkRateSupported       int           `json:"MaxUplinkRateSupported"`
	MaxUplinkRateReached         int           `json:"MaxUplinkRateReached"`
	MaxTxSpatialStreamsSupported int           `json:"MaxTxSpatialStreamsSupported"`
	MaxRxSpatialStreamsSupported int           `json:"MaxRxSpatialStreamsSupported"`
	Index                        string        `json:"Index"`
	Names                        []Names       `json:"Names"`
	DeviceTypes                  []DeviceTypes `json:"DeviceTypes"`
	Bdd                          Bdd           `json:"BDD"`
	ModelNames                   []any         `json:"ModelNames,omitempty"`
	OperatingSystems             []any         `json:"OperatingSystems,omitempty"`
	SoftwareVersions             []any         `json:"SoftwareVersions,omitempty"`
	Manufacturers                []any         `json:"Manufacturers,omitempty"`
	DeviceCategories             []any         `json:"DeviceCategories,omitempty"`
	CatalogIDs                   []any         `json:"CatalogIDs,omitempty"`
	IPv4Address                  []IPv4Address `json:"IPv4Address"`
	IPv6Address                  []any         `json:"IPv6Address"`
	Locations                    []any         `json:"Locations"`
	Groups                       []any         `json:"Groups"`
	Priority                     Priority      `json:"Priority"`
	Security                     Security      `json:"Security"`
	SSWSta                       SSWSta        `json:"SSWSta"`
	UserAgents                   []any         `json:"UserAgents"`
	WANAccess                    WANAccess     `json:"WANAccess"`
	MDNSService                  []MDNSService `json:"mDNSService,omitempty"`
	MDNSRecord                   []any         `json:"mDNSRecord,omitempty"`
}
type AccessPoint struct {
	Key                              string    `json:"Key"`
	DiscoverySource                  string    `json:"DiscoverySource"`
	Name                             string    `json:"Name"`
	DeviceType                       string    `json:"DeviceType"`
	Active                           bool      `json:"Active"`
	Tags                             string    `json:"Tags"`
	FirstSeen                        time.Time `json:"FirstSeen"`
	LastConnection                   time.Time `json:"LastConnection"`
	LastChanged                      time.Time `json:"LastChanged"`
	Master                           string    `json:"Master"`
	PortState                        string    `json:"PortState"`
	NetDevName                       string    `json:"NetDevName"`
	NetDevIndex                      int       `json:"NetDevIndex"`
	Ssid                             string    `json:"SSID,omitempty"`
	Bssid                            string    `json:"BSSID,omitempty"`
	OperatingFrequencyBand           string    `json:"OperatingFrequencyBand,omitempty"`
	OperatingStandards               string    `json:"OperatingStandards,omitempty"`
	Channel                          int       `json:"Channel,omitempty"`
	CurrentOperatingChannelBandwidth string    `json:"CurrentOperatingChannelBandwidth,omitempty"`
	OperatingClass                   int       `json:"OperatingClass,omitempty"`
	Enabled                          bool      `json:"Enabled,omitempty"`
	EssIdentifier                    string    `json:"EssIdentifier,omitempty"`
	Index                            string    `json:"Index"`
	Names                            []Names   `json:"Names"`
	DeviceTypes                      []any     `json:"DeviceTypes"`
	Wps                              Wps       `json:"WPS,omitempty"`
	Devices                          []Device  `json:"Children,omitempty"`
	MaxBitRateSupported              int       `json:"MaxBitRateSupported,omitempty"`
	CurrentBitRate                   int       `json:"CurrentBitRate,omitempty"`
	CurrentDuplexMode                string    `json:"CurrentDuplexMode,omitempty"`
}
type Status struct {
	Key                    string        `json:"Key"`
	DiscoverySource        string        `json:"DiscoverySource"`
	Name                   string        `json:"Name"`
	DeviceType             string        `json:"DeviceType"`
	Active                 bool          `json:"Active"`
	Tags                   string        `json:"Tags"`
	FirstSeen              time.Time     `json:"FirstSeen"`
	LastConnection         time.Time     `json:"LastConnection"`
	LastChanged            time.Time     `json:"LastChanged"`
	Master                 string        `json:"Master"`
	PortState              string        `json:"PortState"`
	PhysAddress            string        `json:"PhysAddress"`
	Layer2Interface        string        `json:"Layer2Interface"`
	InterfaceName          string        `json:"InterfaceName"`
	MACVendor              string        `json:"MACVendor"`
	NetDevName             string        `json:"NetDevName"`
	NetDevIndex            int           `json:"NetDevIndex"`
	IPAddress              string        `json:"IPAddress"`
	IPAddressSource        string        `json:"IPAddressSource"`
	DHCPv4ServerPool       string        `json:"DHCPv4ServerPool"`
	DHCPv4ServerEnable     bool          `json:"DHCPv4ServerEnable"`
	DHCPv4ServerMinAddress string        `json:"DHCPv4ServerMinAddress"`
	DHCPv4ServerMaxAddress string        `json:"DHCPv4ServerMaxAddress"`
	DHCPv4ServerNetmask    string        `json:"DHCPv4ServerNetmask"`
	DHCPv4DomainName       string        `json:"DHCPv4DomainName"`
	Index                  string        `json:"Index"`
	Names                  []Names       `json:"Names"`
	DeviceTypes            []any         `json:"DeviceTypes"`
	IPv4Address            []IPv4Address `json:"IPv4Address"`
	IPv6Address            []IPv6Address `json:"IPv6Address"`
	AccessPoints           []AccessPoint `json:"Children"`
}
