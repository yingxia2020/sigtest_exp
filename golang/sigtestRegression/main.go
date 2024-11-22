package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func testPWJ() {
	// Define the executable and its parameters
	executable := "C:\\Users\\yingxia\\sandbox\\applications.validation.compliance.io-standards.sigtest-phoenix-code-gen6\\SigTestGUI\\bin\\Debug\\SigTest.exe"
	//executable := "C:\\Program Files\\SigTest 6.1.07\\SigTest.exe"

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
}

func testRjCal() {
	// Define the executable and its parameters
	executable := "C:\\Users\\yingxia\\sandbox\\applications.validation.compliance.io-standards.sigtest-phoenix-code-gen6\\SigTestGUI\\bin\\Debug\\SigTest.exe"
	//executable := "C:\\Program Files\\SigTest 6.1.07\\SigTest.exe"

	// Second test: Rj test
	// .\SigTest.exe PCIe 6_0 CEM_RjCal /t Rj_Cal /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\4_Sample_Gen5_Data_jacob\5_0_RXCal\Rj_SjCal\Keysight_Rj_Cal_Waveforms_0.27ps_000.bin"
	params := []string{"PCIe", "6_0", "CEM_RjCal", "/t", "Rj_Cal", "/s",
		"C:\\SigTest\\regression_tests\\gen6_rjsj\\Keysight_Rj_Cal_Waveforms_0.27ps_000.bin"}

	// Create the command
	cmd := exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	fmt.Println("Rj Cal flow regression test started...")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	// Get the output
	output := outBuffer.String()

	if strings.Contains(output, "31.25009E-12,0.2149E-12,2.19168E-12,0.00000E0,0E-12,Rj_Cal,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("Rj Cal flow regression test passed!")
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

	if strings.Contains(output, "31.25011E-12,0.15458E-12,1.69844E-12,3.27519E-3,0.074672E-12,Rj_Cal,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("Rj Cal with SN flow regression test passed!")
	} else {
		fmt.Println("Test failed!")
		fmt.Printf("Output: %s\n", output)
	}
}

func testSjCal() {
	// Define the executable and its parameters
	executable := "C:\\Users\\yingxia\\sandbox\\applications.validation.compliance.io-standards.sigtest-phoenix-code-gen6\\SigTestGUI\\bin\\Debug\\SigTest.exe"
	//executable := "C:\\Program Files\\SigTest 6.1.07\\SigTest.exe"

	// Third test: Sj test
	// .\SigTest.exe PCIe 6_0 CEM_SjCal /t Sj_Cal /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\4_Sample_Gen5_Data_jacob\5_0_RXCal\Rj_SjCal\Keysight_Sj_Cal_Waveforms_0.4ps_000.bin"
	params := []string{"PCIe", "6_0", "CEM_SjCal", "/t", "Sj_Cal", "/s",
		"C:\\SigTest\\regression_tests\\gen6_rjsj\\Keysight_Sj_Cal_Waveforms_0.4ps_000.bin"}

	// Create the command
	cmd := exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	fmt.Println("Sj Cal flow regression test started...")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	// Get the output
	output := outBuffer.String()

	if strings.Contains(output, "31.25009E-12,0.21071E-12,3.65E-12,Sj_Cal,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("Sj Cal flow regression test passed!")
	} else {
		fmt.Println("Test failed!")
		fmt.Printf("Output: %s\n", output)
	}
}

func testCompliance() {
	// Define the executable and its parameters
	executable := "C:\\Users\\yingxia\\sandbox\\applications.validation.compliance.io-standards.sigtest-phoenix-code-gen6\\SigTestGUI\\bin\\Debug\\SigTest.exe"
	//executable := "C:\\Program Files\\SigTest 6.1.07\\SigTest.exe"

	// Fourth test: Compliance test
	// .\SigTest.exe PCIe 6_0 Base_TX_w_CompliancePattn /t EIEOS_FS_P10 /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_compliance\2023ww11_BERT2SCOPE_SNDRNonAveragedWavefrom.bin"
	params := []string{"PCIe", "6_0", "Base_TX_w_CompliancePattn", "/t", "EIEOS_FS_P10", "/s",
		"C:\\SigTest\\regression_tests\\gen6_compliance\\2023ww11_BERT2SCOPE_SNDRNonAveragedWavefrom.bin"}

	// Create the command
	cmd := exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	fmt.Println("Compliance flow regression test started...")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	// Get the output
	output := outBuffer.String()

	// results with brickwall filter
	if strings.Contains(output, "0.754726,-1.74553,0.753927,EIEOS_FS_P10, , ,PASS") &&
		// results without brickwall filter
		//if strings.Contains(output, "0.754726,-1.707674,0.753927,EIEOS_FS_P10, , ,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("Compliance flow regression test passed!")
	} else {
		fmt.Println("Test failed!")
		fmt.Printf("Output: %s\n", output)
	}
}

func testPreset() {
	// Define the executable and its parameters
	executable := "C:\\Users\\yingxia\\sandbox\\applications.validation.compliance.io-standards.sigtest-phoenix-code-gen6\\SigTestGUI\\bin\\Debug\\SigTest.exe"
	//executable := "C:\\Program Files\\SigTest 6.1.07\\SigTest.exe"

	// Fifth test: Preset test
	// .\SigTest.exe PCIe 6_0 Preset_TestAC_SingleRun /t No_CTLE /fs /d "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_preset" /wrefd  NoChannel_Q0_d_000.wfm /wtstd NoChannel_Q2_d_000.wfm
	params := []string{"PCIe", "6_0", "Preset_TestAC_SingleRun", "/t", "No_CTLE", "/fs",
		"/d", "C:\\SigTest\\regression_tests\\gen6_preset",
		"/wrefd", "NoChannel_Q0_d_000.wfm",
		"/wtstd", "NoChannel_Q2_d_000.wfm"}

	// Create the command
	cmd := exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	fmt.Println("Preset AC single flow regression test started...")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	// Get the output
	output := outBuffer.String()

	if strings.Contains(output, "0.00795,-0.17256,0.80573,0.00476,0.17066,0.99099,63,0.5008,-10.87134,50.76074,0.29979,-0.21648,3.71848,0.12894,3.63466,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("Preset AC single flow regression test passed!")
	} else {
		fmt.Println("Test failed!")
		fmt.Printf("Output: %s\n", output)
	}
}

func testTIE() {
	// Define the executable and its parameters
	executable := "C:\\Users\\yingxia\\sandbox\\applications.validation.compliance.io-standards.sigtest-phoenix-code-gen6\\SigTestGUI\\bin\\Debug\\SigTest.exe"
	//executable := "C:\\Program Files\\SigTest 6.1.07\\SigTest.exe"

	// Sixth test: TIE test
	// .\SigTest.exe PCIe 6_0 Base_TX_w_52UIJitterPattn /t 5dB_CTLE /wd "C:\Sigtest\regression_tests\gen6_tie\tek_signal" /nd "C:\Sigtest\regression_tests\gen6_tie\tek_noise"
	params := []string{"PCIe", "6_0", "Base_TX_w_52UIJitterPattn", "/t", "5dB_CTLE",
		"/wf", "C:\\SigTest\\regression_tests\\gen6_tie\\tek_signal\\2023ww24_BERT2Scope_2dB_S00_d_52UIjitter_noEq_60mV_10MUI_2.wfm",
		"/nf", "C:\\SigTest\\regression_tests\\gen6_tie\\tek_noise\\2023ww24_BERT2Scope_2dB_N00_d_noiseWaveform_noEq_60mV_10MUI.wfm"}

	// Create the command
	cmd := exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	fmt.Println("TIE flow regression test started...")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	// Get the output
	output := outBuffer.String()

	if strings.Contains(output, "3.617754E-12,0.718955E-12,0.223485E-12,0.294826E-12,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("TIE flow regression test passed!")
	} else {
		fmt.Println("Test failed!")
		fmt.Printf("Output: %s\n", output)
	}
}

func testSNDR() {
	// Define the executable and its parameters
	executable := "C:\\Users\\yingxia\\sandbox\\applications.validation.compliance.io-standards.sigtest-phoenix-code-gen6\\SigTestGUI\\bin\\Debug\\SigTest.exe"
	//executable := "C:\\Program Files\\SigTest 6.1.07\\SigTest.exe"

	// Eighth test: SNDR test
	// .\SigTest.exe PCIe 6_0 SNDR_RLM /t SNDR /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_sndr\ScopeA_LatestFW_UXI64G_CH24_run1.bin" /wpat "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_sndr\PCIE6_compliance_137blocks_lane0_Rev6p0p1.pat"
	params := []string{"PCIe", "6_0", "SNDR_RLM", "/t", "SNDR+RLM", "/s",
		"C:\\SigTest\\regression_tests\\SNDR\\ScopeA_LatestFW_UXI64G_CH24_run1.bin", "/pat",
		"lane0"}

	// Create the command
	cmd := exec.Command(executable, params...)

	// Create a buffer to capture the standard output and standard error
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	fmt.Println("SNDR_RLM flow regression test started...")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Standard Error: %s\n", errBuffer.String())
		return
	}

	// Get the output
	output := outBuffer.String()

	if strings.Contains(output, "32,250,174.583115,34.237343,2.496148,2.293017,N/A,0.965876,0.992232,SNDR+RLM,lane0,PASS") &&
		!strings.HasSuffix(output,",") {
		fmt.Println("SNDR_RLM flow regression test passed!")
	} else {
		fmt.Println("Test failed!")
		fmt.Printf("Output: %s\n", output)
	}
}

// Map test case names to functions
var testCases = map[string]func(){
	"1": testPWJ,
	"2": testRjCal,
	"3": testSjCal,
	"4": testCompliance,
	"5": testPreset,
	"6": testTIE,
	"7": testSNDR,
}

// Ordered list of test case keys
var testCaseOrder = []string{"1", "2", "3", "4", "5", "6", "7"}

func main() {
	//// Define the executable and its parameters
	//executable := "C:\\Users\\yingxia\\sandbox\\applications.validation.compliance.io-standards.sigtest-phoenix-code-gen6\\SigTestGUI\\bin\\Debug\\SigTest.exe"
	////executable := "C:\\Program Files\\SigTest 6.1.07\\SigTest.exe"
	//
	//// First test: PWJ
	//// .\SigTest.exe PCIe 6_0 Base_TX_w_ClockPattn /t No_CTLE /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_new_pwj\2024ww03_Toggle_50GBW_33GBT_50mV000.wfm" /nf "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_new_pwj\2024ww03_noise_50GBW_33GBT_50mV000.wfm"
	//params := []string{"PCIe", "6_0", "Base_TX_w_ClockPattn", "/t", "No_CTLE", "/s",
	//	"C:\\SigTest\\regression_tests\\gen6_pwj\\2024ww03_Toggle_50GBW_33GBT_50mV000.wfm", "/nf",
	//	"C:\\SigTest\\regression_tests\\gen6_pwj\\2024ww03_noise_50GBW_33GBT_50mV000.wfm"}
	//
	//// Create the command
	//cmd := exec.Command(executable, params...)
	//
	//// Create a buffer to capture the standard output and standard error
	//var outBuffer bytes.Buffer
	//var errBuffer bytes.Buffer
	//cmd.Stdout = &outBuffer
	//cmd.Stderr = &errBuffer
	//
	//// Run the command
	//fmt.Println("PWJ flow regression test started...")
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Printf("Error: %s\n", err)
	//	fmt.Printf("Standard Error: %s\n", errBuffer.String())
	//	return
	//}
	//
	//// Get the output
	//output := outBuffer.String()
	//
	//if strings.Contains(output, "0.165331E-12,0.082666E-12,1.904275E-12,0.203677E-12,0.166368E-12,No_CTLE,0,3.12501E-11,3.27519E-3,PASS") &&
	//	!strings.HasSuffix(output,",") {
	//	fmt.Println("PWJ flow regression test passed!")
	//} else {
	//	fmt.Println("Test failed!")
	//	fmt.Printf("Output: %s\n", output)
	//}
	//
	//// Second test: Rj test
	//// .\SigTest.exe PCIe 6_0 CEM_RjCal /t Rj_Cal /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\4_Sample_Gen5_Data_jacob\5_0_RXCal\Rj_SjCal\Keysight_Rj_Cal_Waveforms_0.27ps_000.bin"
	//params = []string{"PCIe", "6_0", "CEM_RjCal", "/t", "Rj_Cal", "/s",
	//	"C:\\SigTest\\regression_tests\\gen6_rjsj\\Keysight_Rj_Cal_Waveforms_0.27ps_000.bin"}
	//
	//// Create the command
	//cmd = exec.Command(executable, params...)
	//
	//// Create a buffer to capture the standard output and standard error
	//outBuffer.Reset()
	//errBuffer.Reset()
	//cmd.Stdout = &outBuffer
	//cmd.Stderr = &errBuffer
	//
	//// Run the command
	//fmt.Println("Rj Cal flow regression test started...")
	//err = cmd.Run()
	//if err != nil {
	//	fmt.Printf("Error: %s\n", err)
	//	fmt.Printf("Standard Error: %s\n", errBuffer.String())
	//	return
	//}
	//
	//// Get the output
	//output = outBuffer.String()
	//
	//if strings.Contains(output, "31.25009E-12,0.2149E-12,2.19168E-12,0.00000E0,0E-12,Rj_Cal,PASS") &&
	//	!strings.HasSuffix(output,",") {
	//	fmt.Println("Rj Cal flow regression test passed!")
	//} else {
	//	fmt.Println("Test failed!")
	//	fmt.Printf("Output: %s\n", output)
	//}
	//
	//// Third test: Sj test
	//// .\SigTest.exe PCIe 6_0 CEM_SjCal /t Sj_Cal /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\4_Sample_Gen5_Data_jacob\5_0_RXCal\Rj_SjCal\Keysight_Sj_Cal_Waveforms_0.4ps_000.bin"
	//params = []string{"PCIe", "6_0", "CEM_SjCal", "/t", "Sj_Cal", "/s",
	//	"C:\\SigTest\\regression_tests\\gen6_rjsj\\Keysight_Sj_Cal_Waveforms_0.4ps_000.bin"}
	//
	//// Create the command
	//cmd = exec.Command(executable, params...)
	//
	//// Create a buffer to capture the standard output and standard error
	//outBuffer.Reset()
	//errBuffer.Reset()
	//cmd.Stdout = &outBuffer
	//cmd.Stderr = &errBuffer
	//
	//// Run the command
	//fmt.Println("Sj Cal flow regression test started...")
	//err = cmd.Run()
	//if err != nil {
	//	fmt.Printf("Error: %s\n", err)
	//	fmt.Printf("Standard Error: %s\n", errBuffer.String())
	//	return
	//}
	//
	//// Get the output
	//output = outBuffer.String()
	//
	//if strings.Contains(output, "31.25009E-12,0.21071E-12,3.65E-12,Sj_Cal,PASS") &&
	//	!strings.HasSuffix(output,",") {
	//	fmt.Println("Sj Cal flow regression test passed!")
	//} else {
	//	fmt.Println("Test failed!")
	//	fmt.Printf("Output: %s\n", output)
	//}
	//
	//// Fourth test: Compliance test
	//// .\SigTest.exe PCIe 6_0 Base_TX_w_CompliancePattn /t EIEOS_FS_P10 /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_compliance\2023ww11_BERT2SCOPE_SNDRNonAveragedWavefrom.bin"
	//params = []string{"PCIe", "6_0", "Base_TX_w_CompliancePattn", "/t", "EIEOS_FS_P10", "/s",
	//	"C:\\SigTest\\regression_tests\\gen6_compliance\\2023ww11_BERT2SCOPE_SNDRNonAveragedWavefrom.bin"}
	//
	//// Create the command
	//cmd = exec.Command(executable, params...)
	//
	//// Create a buffer to capture the standard output and standard error
	//outBuffer.Reset()
	//errBuffer.Reset()
	//cmd.Stdout = &outBuffer
	//cmd.Stderr = &errBuffer
	//
	//// Run the command
	//fmt.Println("Compliance flow regression test started...")
	//err = cmd.Run()
	//if err != nil {
	//	fmt.Printf("Error: %s\n", err)
	//	fmt.Printf("Standard Error: %s\n", errBuffer.String())
	//	return
	//}
	//
	//// Get the output
	//output = outBuffer.String()
	//
	//// results with brickwall filter
	//if strings.Contains(output, "0.754726,-1.74553,0.753927,EIEOS_FS_P10, , ,PASS") &&
	//// results without brickwall filter
	////if strings.Contains(output, "0.754726,-1.707674,0.753927,EIEOS_FS_P10, , ,PASS") &&
	//	!strings.HasSuffix(output,",") {
	//	fmt.Println("Compliance flow regression test passed!")
	//} else {
	//	fmt.Println("Test failed!")
	//	fmt.Printf("Output: %s\n", output)
	//}
	//
	//// Fifth test: Preset test
	//// .\SigTest.exe PCIe 6_0 Preset_TestAC_SingleRun /t No_CTLE /fs /d "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_preset" /wrefd  NoChannel_Q0_d_000.wfm /wtstd NoChannel_Q2_d_000.wfm
	//params = []string{"PCIe", "6_0", "Preset_TestAC_SingleRun", "/t", "No_CTLE", "/fs",
	//	"/d", "C:\\SigTest\\regression_tests\\gen6_preset",
	//	"/wrefd", "NoChannel_Q0_d_000.wfm",
	//	"/wtstd", "NoChannel_Q2_d_000.wfm"}
	//
	//// Create the command
	//cmd = exec.Command(executable, params...)
	//
	//// Create a buffer to capture the standard output and standard error
	//outBuffer.Reset()
	//errBuffer.Reset()
	//cmd.Stdout = &outBuffer
	//cmd.Stderr = &errBuffer
	//
	//// Run the command
	//fmt.Println("Preset AC single flow regression test started...")
	//err = cmd.Run()
	//if err != nil {
	//	fmt.Printf("Error: %s\n", err)
	//	fmt.Printf("Standard Error: %s\n", errBuffer.String())
	//	return
	//}
	//
	//// Get the output
	//output = outBuffer.String()
	//
	//if strings.Contains(output, "0.00795,-0.17256,0.80573,0.00476,0.17066,0.99099,63,0.5008,-10.87134,50.76074,0.29979,-0.21648,3.71848,0.12894,3.63466,PASS") &&
	//	!strings.HasSuffix(output,",") {
	//	fmt.Println("Preset AC single flow regression test passed!")
	//} else {
	//	fmt.Println("Test failed!")
	//	fmt.Printf("Output: %s\n", output)
	//}
	//
	//// Sixth test: TIE test
	//// .\SigTest.exe PCIe 6_0 Base_TX_w_52UIJitterPattn /t 5dB_CTLE /wd "C:\Sigtest\regression_tests\gen6_tie\tek_signal" /nd "C:\Sigtest\regression_tests\gen6_tie\tek_noise"
	//params = []string{"PCIe", "6_0", "Base_TX_w_52UIJitterPattn", "/t", "5dB_CTLE",
	//	"/wf", "C:\\SigTest\\regression_tests\\gen6_tie\\tek_signal\\2023ww24_BERT2Scope_2dB_S00_d_52UIjitter_noEq_60mV_10MUI_2.wfm",
	//	"/nf", "C:\\SigTest\\regression_tests\\gen6_tie\\tek_noise\\2023ww24_BERT2Scope_2dB_N00_d_noiseWaveform_noEq_60mV_10MUI.wfm"}
	//
	//// Create the command
	//cmd = exec.Command(executable, params...)
	//
	//// Create a buffer to capture the standard output and standard error
	//outBuffer.Reset()
	//errBuffer.Reset()
	//cmd.Stdout = &outBuffer
	//cmd.Stderr = &errBuffer
	//
	//// Run the command
	//fmt.Println("TIE flow regression test started...")
	//err = cmd.Run()
	//if err != nil {
	//	fmt.Printf("Error: %s\n", err)
	//	fmt.Printf("Standard Error: %s\n", errBuffer.String())
	//	return
	//}
	//
	//// Get the output
	//output = outBuffer.String()
	//
	//if strings.Contains(output, "3.617754E-12,0.718955E-12,0.223485E-12,0.294826E-12,PASS") &&
	//	!strings.HasSuffix(output,",") {
	//	fmt.Println("TIE flow regression test passed!")
	//} else {
	//	fmt.Println("Test failed!")
	//	fmt.Printf("Output: %s\n", output)
	//}
	//
	//// Seventh test: Rj test with scope noise
	//// .\SigTest.exe PCIe 6_0 CEM_RjCal /t Rj_Cal /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_new_pwj\2024ww03_Toggle_50GBW_33GBT_50mV000.wfm" /nf "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_new_pwj\2024ww03_noise_50GBW_33GBT_50mV000.wfm"
	//params = []string{"PCIe", "6_0", "CEM_RjCal", "/t", "Rj_Cal", "/s",
	//	"C:\\SigTest\\regression_tests\\gen6_pwj\\2024ww03_Toggle_50GBW_33GBT_50mV000.wfm", "/nf",
	//	"C:\\SigTest\\regression_tests\\gen6_pwj\\2024ww03_noise_50GBW_33GBT_50mV000.wfm"}
	//
	//// Create the command
	//cmd = exec.Command(executable, params...)
	//
	//// Create a buffer to capture the standard output and standard error
	//outBuffer.Reset()
	//errBuffer.Reset()
	//cmd.Stdout = &outBuffer
	//cmd.Stderr = &errBuffer
	//
	//// Run the command
	//fmt.Println("Rj Cal with SN flow regression test started...")
	//err = cmd.Run()
	//if err != nil {
	//	fmt.Printf("Error: %s\n", err)
	//	fmt.Printf("Standard Error: %s\n", errBuffer.String())
	//	return
	//}
	//
	//// Get the output
	//output = outBuffer.String()
	//
	//if strings.Contains(output, "31.25011E-12,0.15458E-12,1.69844E-12,3.27519E-3,0.074672E-12,Rj_Cal,PASS") &&
	//	!strings.HasSuffix(output,",") {
	//	fmt.Println("Rj Cal with SN flow regression test passed!")
	//} else {
	//	fmt.Println("Test failed!")
	//	fmt.Printf("Output: %s\n", output)
	//}
	//
	//// Eighth test: SNDR test
	//// .\SigTest.exe PCIe 6_0 SNDR_RLM /t SNDR /s "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_sndr\ScopeA_LatestFW_UXI64G_CH24_run1.bin" /wpat "C:\Users\yingxia\Downloads\SigTest\wave_forms\gen6_sndr\PCIE6_compliance_137blocks_lane0_Rev6p0p1.pat"
	//params = []string{"PCIe", "6_0", "SNDR_RLM", "/t", "SNDR+RLM", "/s",
	//	"C:\\SigTest\\regression_tests\\SNDR\\ScopeA_LatestFW_UXI64G_CH24_run1.bin", "/pat",
	//	"lane0"}
	//
	//// Create the command
	//cmd = exec.Command(executable, params...)
	//
	//// Create a buffer to capture the standard output and standard error
	//outBuffer.Reset()
	//errBuffer.Reset()
	//cmd.Stdout = &outBuffer
	//cmd.Stderr = &errBuffer
	//
	//// Run the command
	//fmt.Println("SNDR_RLM flow regression test started...")
	//err = cmd.Run()
	//if err != nil {
	//	fmt.Printf("Error: %s\n", err)
	//	fmt.Printf("Standard Error: %s\n", errBuffer.String())
	//	return
	//}
	//
	//// Get the output
	//output = outBuffer.String()
	//
	//if strings.Contains(output, "32,250,174.583115,34.237343,2.496148,2.293017,N/A,0.965876,0.992232,SNDR+RLM,lane0,PASS") &&
	//	!strings.HasSuffix(output,",") {
	//	fmt.Println("SNDR_RLM flow regression test passed!")
	//} else {
	//	fmt.Println("Test failed!")
	//	fmt.Printf("Output: %s\n", output)
	//}

	// Print the output
	//fmt.Printf("Output: %s\n", output)

	for {
		fmt.Println("SigTest Regression Test Runner")
		fmt.Println("==============================")
		fmt.Println("1. Test Case 1: testPWJ")
		fmt.Println("2. Test Case 2: testRjCal")
		fmt.Println("3. Test Case 3: testSjCal")
		fmt.Println("4. Test Case 4: testCompliance")
		fmt.Println("5. Test Case 5: testPreset")
		fmt.Println("6. Test Case 6: testTIE")
		fmt.Println("7. Test Case 7: testSNDR")
		fmt.Println("8. Run All Tests")
		fmt.Println("0. Exit")
		fmt.Println("Enter the number of the test case to run, '8' to run all, '0' to exit:")

		// Read user input
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		// Exit condition
		if choice == "0" {
			fmt.Println("Exiting SigTest Regression Test Runner. Goodbye!")
			break
		}

		// Execute based on user choice
		switch choice {
		case "1", "2", "3", "4", "5", "6", "7":
			fmt.Println("\n--- Starting Selected Test ---")
			testCases[choice]()
		case "8":
			fmt.Println("\n--- Starting All Tests ---")
			for _, key := range testCaseOrder {
				fmt.Printf("\nRunning Test Case %s:\n", key)
				testCases[key]()
			}
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
