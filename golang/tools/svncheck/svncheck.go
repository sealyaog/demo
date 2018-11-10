package main

import(
    "fmt"
    "os"
    "flag"
    //"os/exec"
    //"io"
    "bufio"
    //"time"
    //"strconv"
    //"sort"
    //s "strings"
    "github.com/ThomasRooney/gexpect"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func svncheck(svnPath string) {
    //c := exec.Command("svn", "ls","-R", svnPath, "|" , "grep" "BUILD")
    //err := c.Run()
    //check(err)
    child, err := gexpect.Spawn("svn ls -R " + svnPath + " | grep BUILD")
    check(err)

    child.Expect("\u7684\u5bc6\u7801")
    child.SendLine("\n")

    child.Expect("\u7528\u6237\u540d")
    fmt.Println("kkk2");
    child.SendLine("yaogang")
    fmt.Println("kkk3");

    child.Expect("\u7684\u5bc6\u7801")
    fmt.Println("kkk4");
    child.SendLine("811115Abc&")
    fmt.Println("kkk5");

    msg, _ := child.ReadLine()
    fmt.Println(msg);
    child.Interact()
    child.Close()
}

func process(inputFile string, outputFile string) {
    file, err := os.Open(inputFile)
    check(err)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        svncheck(scanner.Text())
    }
}

func ParseOpt(input *string, output *string) {
    flag.StringVar(input, "i", " ", "inputFile")
    flag.StringVar(output, "o", " ", "outputFile")
    flag.Parse()
    fmt.Println("input: ", *input);
    fmt.Println("output: ", *output);
}

func main() {
    var inputFile string;
    var outputFile string;
    ParseOpt(&inputFile, &outputFile);
    process(inputFile, outputFile);
}
