package main

type ReliabilityTestResult struct {
	LogsSent        int `json:"logsSent"`
	LogsReceived    int `json:"logsReceived"`
	BadLogsReceived int `json:"badLogsReceived"`
}
