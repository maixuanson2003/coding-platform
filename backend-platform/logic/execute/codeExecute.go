package execute

import (
	"fmt"
	"io"
	"lietcode/logic/constant"
	"log"
	"strings"

	"os"
	"os/exec"
	"path/filepath"
)

type CodeExecuteWorker struct {
}
type CodeExecuteConfig struct {
	Lang     string
	Code     string
	FileName string
}

func (code *CodeExecuteWorker) SetUpToWorker(config CodeExecuteConfig) (string, error) {
	folder := constant.WorkerFolderPath[config.Lang]
	domain := constant.DomainExtension[config.Lang]
	fileName := fmt.Sprintf("%s.%s", config.FileName, domain)
	filePadding := filepath.Join(folder, fileName)
	file, errCreateFolder := os.Create(filePadding)
	if errCreateFolder != nil {
		fmt.Println("Error creating file:", errCreateFolder)
		return "", errCreateFolder
	}
	defer file.Close()
	injectFunc, ok := InjectStruct[config.Lang]
	var codeCheck string
	if ok {
		codeCheck = injectFunc(config.Code)

	}
	_, errWriteString := file.WriteString(codeCheck)
	if errWriteString != nil {
		fmt.Println("Error writing code:", errWriteString)
		return "", errWriteString
	}
	if config.Lang == "cpp" {
		cmd := exec.Command("clang-format", "-i", filePadding)
		cmd.Run()

	} else if config.Lang == "java" {
		log.Print("run")
		cmd := exec.Command("google-java-format", "-i", filePadding)
		log.Print(cmd.Output())
		cmd.Run()

	} else if config.Lang == "python" {
		cmd := exec.Command("black", filePadding)
		cmd.Run()

	} else if config.Lang == "js" {
		cmd := exec.Command("prettier", "--write", filePadding)
		cmd.Run()
	}

	fmt.Println("✅ File created successfully:", file.Name())
	return fileName, nil

}
func (code *CodeExecuteWorker) ExecuteCppCode(inputs []string, fileName string) (map[uint]map[string]interface{}, error) {
	commandExcuteCode := constant.CppTemplate
	template := strings.ReplaceAll(commandExcuteCode, "{{nameFile}}", fileName)
	outPuts := map[uint]map[string]interface{}{}

	for index, input := range inputs {

		script := strings.ReplaceAll(template, "{{input}}", input)

		cmd := exec.Command("docker", "exec", "-i", "env_cpp", "bash", "-c", script)
		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		if err := cmd.Start(); err != nil {
			return nil, err
		}

		stdoutBytes, _ := io.ReadAll(stdout)
		stderrBytes, _ := io.ReadAll(stderr)

		err := cmd.Wait()

		if err != nil {
			outPuts[uint(index)+1] = map[string]interface{}{
				"error": strings.TrimSpace(string(stderrBytes)),
			}
			continue
		}

		raw := strings.TrimSpace(string(stdoutBytes))
		lines := strings.Split(raw, "\n")

		if len(lines) < 2 {
			outPuts[uint(index)+1] = map[string]interface{}{
				"error": "Invalid output format",
			}
			continue
		}

		timeMs := strings.TrimPrefix(lines[len(lines)-2], "TIME_MS=")
		memoryKb := strings.TrimPrefix(lines[len(lines)-1], "MEMORY_KB=")

		outputs := lines[:len(lines)-2]

		joinedOutput := strings.TrimSpace(strings.Join(outputs, "\n"))

		outPuts[uint(index)+1] = map[string]interface{}{
			"output":    joinedOutput,
			"time_ms":   timeMs,
			"memory_kb": memoryKb,
		}
	}

	return outPuts, nil
}
func (code *CodeExecuteWorker) ExecuteJavaCode(inputs []string, fileName string) (map[uint]map[string]interface{}, error) {
	template := constant.JavaTemplate
	template = strings.ReplaceAll(template, "{{nameFile}}", fileName)

	outPuts := map[uint]map[string]interface{}{}

	for index, input := range inputs {

		script := strings.ReplaceAll(template, "{{input}}", input)

		cmd := exec.Command("docker", "exec", "-i", "env_java", "bash", "-c", script)
		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		if err := cmd.Start(); err != nil {
			return nil, err
		}

		stdoutBytes, _ := io.ReadAll(stdout)
		stderrBytes, _ := io.ReadAll(stderr)

		err := cmd.Wait()

		if err != nil {
			outPuts[uint(index)+1] = map[string]interface{}{
				"error": strings.TrimSpace(string(stderrBytes)),
			}
			continue
		}

		raw := strings.TrimSpace(string(stdoutBytes))
		lines := strings.Split(raw, "\n")

		if len(lines) < 2 {
			outPuts[uint(index)+1] = map[string]interface{}{
				"error": "Invalid output format",
			}
			continue
		}

		timeMs := strings.TrimPrefix(lines[len(lines)-2], "TIME_MS=")
		memoryKb := strings.TrimPrefix(lines[len(lines)-1], "MEMORY_KB=")

		outputs := lines[:len(lines)-2]

		joinedOutput := strings.TrimSpace(strings.Join(outputs, "\n"))

		outPuts[uint(index)+1] = map[string]interface{}{
			"output":    joinedOutput,
			"time_ms":   timeMs,
			"memory_kb": memoryKb,
		}
	}

	return outPuts, nil
}
func (code *CodeExecuteWorker) ExecutePythonCode(inputs []string, fileName string) (map[uint]map[string]interface{}, error) {
	template := constant.PythonTemplate
	template = strings.ReplaceAll(template, "{{nameFile}}", fileName)
	log.Print(fileName)

	outPuts := map[uint]map[string]interface{}{}

	for index, input := range inputs {

		script := strings.ReplaceAll(template, "{{input}}", input)

		cmd := exec.Command("docker", "exec", "-i", "env_python", "bash", "-c", script)
		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		if err := cmd.Start(); err != nil {
			return nil, err
		}

		stdoutBytes, _ := io.ReadAll(stdout)
		stderrBytes, _ := io.ReadAll(stderr)

		err := cmd.Wait()

		if err != nil {
			outPuts[uint(index)+1] = map[string]interface{}{
				"error": strings.TrimSpace(string(stderrBytes)),
			}
			continue
		}

		raw := strings.TrimSpace(string(stdoutBytes))
		lines := strings.Split(raw, "\n")

		if len(lines) < 2 {
			outPuts[uint(index)+1] = map[string]interface{}{
				"error": "Invalid output format",
			}
			continue
		}

		timeMs := strings.TrimPrefix(lines[len(lines)-2], "TIME_MS=")
		memoryKb := strings.TrimPrefix(lines[len(lines)-1], "MEMORY_KB=")

		outputs := lines[:len(lines)-2]

		joinedOutput := strings.TrimSpace(strings.Join(outputs, "\n"))

		outPuts[uint(index)+1] = map[string]interface{}{
			"output":    joinedOutput,
			"time_ms":   timeMs,
			"memory_kb": memoryKb,
		}
	}

	return outPuts, nil
}
func (code *CodeExecuteWorker) ExecuteJavaScriptCode(inputs []string, fileName string) (map[uint]map[string]interface{}, error) {
	template := constant.JsTemplate
	template = strings.ReplaceAll(template, "{{nameFile}}", fileName)

	outPuts := map[uint]map[string]interface{}{}

	for index, input := range inputs {

		script := strings.ReplaceAll(template, "{{input}}", input)

		cmd := exec.Command("docker", "exec", "-i", "env_nodejs", "bash", "-c", script)
		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		if err := cmd.Start(); err != nil {
			return nil, err
		}

		stdoutBytes, _ := io.ReadAll(stdout)
		stderrBytes, _ := io.ReadAll(stderr)

		err := cmd.Wait()

		if err != nil {
			outPuts[uint(index)+1] = map[string]interface{}{
				"error": strings.TrimSpace(string(stderrBytes)),
			}
			continue
		}

		raw := strings.TrimSpace(string(stdoutBytes))
		lines := strings.Split(raw, "\n")

		if len(lines) < 2 {
			outPuts[uint(index)+1] = map[string]interface{}{
				"error": "Invalid output format",
			}
			continue
		}

		timeMs := strings.TrimPrefix(lines[len(lines)-2], "TIME_MS=")
		memoryKb := strings.TrimPrefix(lines[len(lines)-1], "MEMORY_KB=")

		outputs := lines[:len(lines)-2]

		joinedOutput := strings.TrimSpace(strings.Join(outputs, "\n"))

		outPuts[uint(index)+1] = map[string]interface{}{
			"output":    joinedOutput,
			"time_ms":   timeMs,
			"memory_kb": memoryKb,
		}
	}

	return outPuts, nil
}
