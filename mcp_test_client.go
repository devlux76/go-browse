package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type MCPRequest struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params,omitempty"`
}

type MCPResponse struct {
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  interface{}     `json:"error"`
}

func main() {
	cmd := exec.Command("go", "run", "./app/main.go")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Failed to get stdin:", err)
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Failed to get stdout:", err)
		return
	}
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to start MCP server:", err)
		return
	}

	reader := bufio.NewReader(stdout)

	send := func(req MCPRequest) error {
		data, _ := json.Marshal(req)
		_, err := stdin.Write(append(data, '\n'))
		return err
	}

	readResp := func() (string, error) {
		line, err := reader.ReadString('\n')
		return line, err
	}

	// 1. Launch Browser
	launchReq := MCPRequest{
		ID:     1,
		Method: "launch_browser",
	}
	fmt.Println("Sending launch_browser request...")
	send(launchReq)
	resp, err := readResp()
	if err != nil && err != io.EOF {
		fmt.Println("Error reading response:", err)
	} else {
		fmt.Println("launch_browser response:", resp)
	}

	// 2. Stat
	statReq := MCPRequest{
		ID:     2,
		Method: "stat",
	}
	fmt.Println("Sending stat request...")
	send(statReq)
	resp, err = readResp()
	if err != nil && err != io.EOF {
		fmt.Println("Error reading response:", err)
	} else {
		fmt.Println("stat response:", resp)
	}

	// 3. List tools
	listReq := MCPRequest{
		ID:     3,
		Method: "list_tools",
	}
	fmt.Println("Sending list_tools request...")
	send(listReq)
	resp, err = readResp()
	if err != nil && err != io.EOF {
		fmt.Println("Error reading response:", err)
	} else {
		fmt.Println("list_tools response:", resp)
	}

	// 4. Navigate
	navReq := MCPRequest{
		ID:     4,
		Method: "navigate",
		Params: map[string]interface{}{"url": "https://theendoftheinternet.com"},
	}
	fmt.Println("Sending navigate request...")
	send(navReq)
	resp, err = readResp()
	if err != nil && err != io.EOF {
		fmt.Println("Error reading response:", err)
	} else {
		fmt.Println("navigate response:", resp)
	}

	stdin.Close()
	cmd.Wait()
}
