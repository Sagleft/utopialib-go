package structs

type SystemInfo struct {
	BuildABI               string `json:"buildAbi"`
	BuildCPUArchitecture   string `json:"buildCpuArchitecture"`
	BuildNumber            string `json:"build_number"`
	CurrentCPUArchitecture string `json:"currentCpuArchitecture"`
	NetCoreRate            int    `json:"netCoreRate"`
	NetworkCores           int    `json:"networkCores"`
	NetworkEnabled         bool   `json:"networkEnabled"`
	NumberOfConnections    int    `json:"numberOfConnections"`
	PacketCacheSize        int    `json:"packetCacheSize"`
	Uptime                 string `json:"uptime"`
}
