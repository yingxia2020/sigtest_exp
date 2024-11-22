package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	ext := ".bin"

	// Define the executable and its parameters
	executable := "C:\\Users\\yingxia\\sandbox\\applications.validation.compliance.io-standards.sigtest-phoenix-code-gen6\\SigTestGUI\\bin\\Release\\SigTest.exe"
	//executable := "C:\\Program Files\\SigTest 6.1.01\\SigTest.exe"

	params := []string{"PCIe", "6_0", "Base_TX_w_52UIJitterPattn", "/t", "Optimize_CTLE", "/m", "2", "/wf",
		"C:\\Users\\yingxia\\Downloads\\52UI\\TID610601185\\SSC_ON\\52UI\\SAMSUNG_TID_610601185_Ln01_52UI_d_256GSps_105mV_xpxdB_SSSCOn_BT_Y_Cap01.bin",
		"/nf", "C:\\Users\\yingxia\\Downloads\\52UI\\TID610601185\\SSC_ON\\noise\\SAMSUNG_TID_610601185_Ln01_NOISE_d_256GSps_105mV_xpxdB_SSSCOff_BT_Y_Cap_.bin",
		"/o", "C:\\Users\\yingxia\\Downloads\\52UI\\TID610601185\\SSC_ON\\52UI\\result_ssc_off.csv"}

	// Create the command
	cmd := exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	oldStr := "Cap01" + ext
	// Run the command
	fmt.Println("TIE flow test for file: " + oldStr)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	var newStr string
	for i := 2; i <= 5; i++ {
		newStr = fmt.Sprintf("Cap%02d", i) + ext
		params[8] = strings.Replace(params[8], oldStr, newStr, 1)
		oldStr = newStr
		cmd = exec.Command(executable, params...)
		fmt.Println("TIE flow test for file: " + oldStr)
		err = cmd.Run()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			fmt.Printf("Standard Error: %s\n", errBuffer.String())
			return
		}
	}
}
