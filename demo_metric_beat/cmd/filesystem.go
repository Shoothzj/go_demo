package cmd

var filesystemCmd = genMetricCmd("filesystem")

func init() {
	rootCmd.AddCommand(filesystemCmd)
}
