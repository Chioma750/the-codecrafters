// CodeCrafters — Operation Gopher Protocol
       // Module: File Pipeline
       // Author: [Ugwu Chioma Ruth]
       // Squad:  [The Gopher's]


// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: Gophers
// ───────────────────────────────────────────
// Input line types:// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: Gophers
// ───────────────────────────────────────────
// Input line types:
// Number of lines: 15 
// Normal report lines                        						
// 1. Lines in ALL CAPS                          					
// 2. Lines in all lowercase                     						
// 3. Lines starting with TODO:                  						       					  					
// 4. Lines with extra leading/trailing spaces 

// Transformation rules (in order):
// 1. Trim all leading and trailing whitespace     					
// 2. Replace TODO: with ✦ ACTION:                 					        				
// 3. Convert ALL CAPS lines to Title Case         					
// 4. Convert all lowercase lines to uppercase    					
// 5. Reverse the words in any line that contains the word REVERSE


// Output format:
// Header: Yes, Exact Text: "Gopher's Sentinel Field Report - Processed"
// Line numbering format : "1."
// Summary block: yes
//  	Fields : 
//		✦ Lines read    :                  						
//		✦ Lines written :                  						
//		✦ Lines removed :                  					
//		✦ Rules applied : [our 5 rules] 
//
//
// Terminal summary fields:
//		✦ Lines read    :                  						
//		✦ Lines written :                  						
//		✦ Lines removed :                  					
//		✦ Rules applied : [our 5 rules]
// ═══════════════════════════════════════════

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Error: ")
        return
    }

    inputFile := os.Args[1]
    outputFile := os.Args[2]

    if inputFile == outputFile {
        fmt.Println("Input and output cannot be the same file.")
        return
    }

    input, err := os.Open(inputFile)
    if err != nil {
        fmt.Println("File not found: input.txt")
        return
    }
    defer input.Close()

    scanner := bufio.NewScanner(input)
    var processedLines []string
    linesRead := 0
    for scanner.Scan() {
        line := scanner.Text()
        line = trimSpaces(line)
        line = replaceTODO(line)
        line = capsToTitle(line)
        line = lowerToUpper(line)
        line = reverseIfNeeded(line)
        processedLines = append(processedLines, line)
        linesRead++
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    if linesRead == 0 {
        fmt.Println("Input file is empty. Nothing to process.")
    }

    output, err := os.Create(outputFile)
    if err != nil {
        fmt.Println("File not found: input.txt")
        return
    }

    fmt.Fprintf(output, "THE GO\n")

    linesWritten := 0
    for i, line := range processedLines {
        fmt.Fprintf(output, "%03d. %s\n", i+1, line)
        linesWritten++
    }

    linesRemoved := linesRead - linesWritten

    fmt.Println("Lines read    :", linesRead)
    fmt.Println("Lines written :", linesWritten)
    fmt.Println("Lines removed :", linesRemoved)
    fmt.Println("Rules applied :", "Trim, TODO→ACTION, ALL CAPS→Title, lowercase→UPPERCASE, REVERSE words")
}

func trimSpaces(line string) string {
    return strings.TrimSpace(line)
}

func replaceTODO(line string) string {
    line = strings.ReplaceAll(line, "todo:", "✦ action:")
    return strings.ReplaceAll(line, "TODO:", "✦ ACTION:")
}

func capsToTitle(line string) string {
    if strings.ToLower(line) != strings.ToUpper(line) && line == strings.ToUpper(line) {
        line = strings.Title(strings.ToLower(line))
    }
    return line
}

func lowerToUpper(line string) string {
    if line == strings.ToLower(line) && strings.ToLower(line) != strings.ToUpper(line) {
        line = strings.ToUpper(line)
    }
    return line
}

func reverseIfNeeded(line string) string {
    if strings.Contains(strings.ToUpper(line), "REVERSE") {
        words := strings.Fields(line)
        for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
            words[i], words[j] = words[j], words[i]
        }
        line = strings.Join(words, " ")
    }
    return line
}



