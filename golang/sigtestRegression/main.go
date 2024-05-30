package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Define the executable and its parameters
	//executable := "C:\\Users\\yingxia\\sandbox\\applications.validation.compliance.io-standards.sigtest-phoenix-code-gen6\\SigTestGUI\\bin\\Release\\SigTest.exe"
	executable := "C:\\Program Files\\SigTest 6.1.01\\SigTest.exe"
	
	// First test: PWJ
	// .\SigTest.exe PCIe 6_0 Base_TX_w_ClockPattn /t No_CTLE /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_new_pwj\2024ww03_Toggle_50GBW_33GBT_50mV000.wfm" /nf "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_new_pwj\2024ww03_noise_50GBW_33GBT_50mV000.wfm"
	params := []string{"PCIe", "6_0", "Base_TX_w_ClockPattn", "/t", "No_CTLE", "/s",
		"C:\\SigTest\\regression_tests\\gen6_pwj\\2024ww03_Toggle_50GBW_33GBT_50mV000.wfm", "/nf",
		"C:\\SigTest\\regression_tests\\gen6_pwj\\2024ww03_noise_50GBW_33GBT_50mV000.wfm"}

	// Create the command
	cmd := exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	fmt.Println("PWJ flow regression test started...")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	// Get the output
	output := outBuffer.String()

	if strings.Contains(output, "0.165331E-12,0.082666E-12,1.904275E-12,0.203677E-12,0.166368E-12,No_CTLE,0,3.12501E-11,3.27519E-3,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("PWJ flow regression test passed!")
	} else {
		fmt.Println("Test failed!")
		fmt.Printf("Output: %s\n", output)
	}

	// Second test: Rj test
	// .\SigTest.exe PCIe 6_0 CEM_RjCal /t Rj_Cal /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\4_Sample_Gen5_Data_jacob\5_0_RXCal\Rj_SjCal\Keysight_Rj_Cal_Waveforms_0.27ps_000.bin"
	params = []string{"PCIe", "6_0", "CEM_RjCal", "/t", "Rj_Cal", "/s",
		"C:\\SigTest\\regression_tests\\gen6_rjsj\\Keysight_Rj_Cal_Waveforms_0.27ps_000.bin"}

	// Create the command
	cmd = exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	outBuffer.Reset()
	errBuffer.Reset()
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	fmt.Println("Rj Cal flow regression test started...")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	// Get the output
	output = outBuffer.String()

	if strings.Contains(output, "31.25009E-12,0.31842E-12,3.69682E-12,0.00000E0,0E-12,Rj_Cal,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("Rj Cal flow regression test passed!")
	} else {
		fmt.Println("Test failed!")
		fmt.Printf("Output: %s\n", output)
	}

	// Third test: Sj test
	// .\SigTest.exe PCIe 6_0 CEM_SjCal /t Sj_Cal /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\4_Sample_Gen5_Data_jacob\5_0_RXCal\Rj_SjCal\Keysight_Sj_Cal_Waveforms_0.4ps_000.bin"
	params = []string{"PCIe", "6_0", "CEM_SjCal", "/t", "Sj_Cal", "/s",
		"C:\\SigTest\\regression_tests\\gen6_rjsj\\Keysight_Sj_Cal_Waveforms_0.4ps_000.bin"}

	// Create the command
	cmd = exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	outBuffer.Reset()
	errBuffer.Reset()
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	fmt.Println("Sj Cal flow regression test started...")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	// Get the output
	output = outBuffer.String()

	if strings.Contains(output, "31.25009E-12,0.40305E-12,4.7E-12,Sj_Cal,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("Sj Cal flow regression test passed!")
	} else {
		fmt.Println("Test failed!")
		fmt.Printf("Output: %s\n", output)
	}

	// Fourth test: Compliance test
	// .\SigTest.exe PCIe 6_0 Base_TX_w_CompliancePattn /t EIEOS_FS_P10 /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_compliance\2023ww11_BERT2SCOPE_SNDRNonAveragedWavefrom.bin"
	params = []string{"PCIe", "6_0", "Base_TX_w_CompliancePattn", "/t", "EIEOS_FS_P10", "/s",
		"C:\\SigTest\\regression_tests\\gen6_compliance\\2023ww11_BERT2SCOPE_SNDRNonAveragedWavefrom.bin"}

	// Create the command
	cmd = exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	outBuffer.Reset()
	errBuffer.Reset()
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	fmt.Println("Compliance flow regression test started...")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	// Get the output
	output = outBuffer.String()

	if strings.Contains(output, "0.754726,-1.687698,0.753927,EIEOS_FS_P10, , ,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("Compliance flow regression test passed!")
	} else {
		fmt.Println("Test failed!")
		fmt.Printf("Output: %s\n", output)
	}

	// Fifth test: Preset test
	// .\SigTest.exe PCIe 6_0 Preset_TestAC_SingleRun /t No_CTLE /fs /d "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_preset" /wrefd  NoChannel_Q0_d_000.wfm /wtstd NoChannel_Q2_d_000.wfm
	params = []string{"PCIe", "6_0", "Preset_TestAC_SingleRun", "/t", "No_CTLE", "/fs",
		"/d", "C:\\SigTest\\regression_tests\\gen6_preset",
		"/wrefd", "NoChannel_Q0_d_000.wfm",
		"/wtstd", "NoChannel_Q2_d_000.wfm"}

	// Create the command
	cmd = exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	outBuffer.Reset()
	errBuffer.Reset()
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	fmt.Println("Preset AC single flow regression test started...")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	// Get the output
	output = outBuffer.String()

	if strings.Contains(output, "0.00795,-0.17256,0.80573,0.00476,0.17066,0.99099,63,0.5008,-10.87134,50.76074,0.29979,-0.21648,3.71848,0.12894,3.63466,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("Preset AC single flow regression test passed!")
	} else {
		fmt.Println("Test failed!")
		fmt.Printf("Output: %s\n", output)
	}

	// Sixth test: TIE test
	// .\SigTest.exe PCIe 6_0 Base_TX_w_52UIJitterPattn /t 5dB_CTLE /wd "C:\Sigtest\regression_tests\gen6_tie\tek_signal" /nd "C:\Sigtest\regression_tests\gen6_tie\tek_noise"
	params = []string{"PCIe", "6_0", "Base_TX_w_52UIJitterPattn", "/t", "5dB_CTLE",
		"/wd", "C:\\SigTest\\regression_tests\\gen6_tie\\tek_signal",
		"/nd", "C:\\SigTest\\regression_tests\\gen6_tie\\tek_noise"}

	// Create the command
	cmd = exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	outBuffer.Reset()
	errBuffer.Reset()
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	fmt.Println("TIE flow regression test started...")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	// Get the output
	output = outBuffer.String()

	if strings.Contains(output, "3.527107E-12,0.753208E-12,0.204759E-12,0.299468E-12,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("TIE flow regression test passed!")
	} else {
		fmt.Println("Test failed!")
		fmt.Printf("Output: %s\n", output)
	}

	// Seventh test: Rj test with scope noise
	// .\SigTest.exe PCIe 6_0 CEM_RjCal /t Rj_Cal /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_new_pwj\2024ww03_Toggle_50GBW_33GBT_50mV000.wfm" /nf "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_new_pwj\2024ww03_noise_50GBW_33GBT_50mV000.wfm"
	params = []string{"PCIe", "6_0", "CEM_RjCal", "/t", "Rj_Cal", "/s",
		"C:\\SigTest\\regression_tests\\gen6_pwj\\2024ww03_Toggle_50GBW_33GBT_50mV000.wfm", "/nf",
		"C:\\SigTest\\regression_tests\\gen6_pwj\\2024ww03_noise_50GBW_33GBT_50mV000.wfm"}

	// Create the command
	cmd = exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	outBuffer.Reset()
	errBuffer.Reset()
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	fmt.Println("Rj Cal with SN flow regression test started...")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	// Get the output
	output = outBuffer.String()

	if strings.Contains(output, "31.25011E-12,0.27459E-12,3.51281E-12,3.27519E-3,0.072913E-12,Rj_Cal,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("Rj Cal with SN flow regression test passed!")
	} else {
		fmt.Println("Test failed!")
		fmt.Printf("Output: %s\n", output)
	}

	// Print the output
	//fmt.Printf("Output: %s\n", output)
}