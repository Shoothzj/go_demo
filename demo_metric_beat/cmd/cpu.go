package cmd

var cpuCmd = genMetricCmd("cpu")

func init() {
	rootCmd.AddCommand(cpuCmd)
}
