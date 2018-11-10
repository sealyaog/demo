package main

import(  
    "fmt"  
    "os"  
    "flag"  
    "io"  
    "bufio"  
    "time"
    "strconv"
    "sort"
    s "strings"
)

type PauseInfo struct {
    grpId uint64 
    isPause bool
    clickTime time.Time
}

type GrpInfo [] *PauseInfo 
type Accountinfo map[uint64] GrpInfo

func (p GrpInfo) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p GrpInfo) Len() int           { return len(p) }
func (p GrpInfo) Less(i, j int) bool { return p[i].grpId < p[j].grpId }


func check(e error) {
    if e != nil {
        panic(e)
    }   
}

var gaccinfo Accountinfo 

const TimeFormat = "2006-01-02 15:04:05"

func printMap() {
    for k, v := range gaccinfo { 
        sort.Sort(v);
        fmt.Printf("account=%v:\n", k)
        for _, v1 := range v {
            fmt.Printf("\t");
            fmt.Printf("grpid=%v ispause=%v time=%v\n", v1.grpId, v1.isPause, v1.clickTime);
        }
    }
}

func readFile(fileName string) {
    f, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    rd := bufio.NewReader(f)

    gaccinfo = Accountinfo{};

    for {
        line, err := rd.ReadString('\n') //read one line
        if err != nil && io.EOF == err {
            break
        }
        check(err)

        //fmt.Println(line)
        split1 := s.Split(s.TrimRight(line, "\n"), "\x03");

        accountId, err := strconv.ParseUint(split1[2], 10, 64)
        check(err)

        grpId, err := strconv.ParseUint(s.Split(split1[0],"\x05")[2], 10, 64);
        check(err)

        isPause := (split1[15] == "1")
        clicktime := s.Split(split1[14],"\x05")[0]

        info := new(PauseInfo)
        info.grpId = grpId
        info.isPause = isPause
        //fmt.Println("time: ", clicktime);

        info.clickTime, err = time.Parse(TimeFormat, clicktime)
        check(err)

        grp := gaccinfo[accountId]
        if(grp == nil) {
            var grpInfo []*PauseInfo 
            grpInfo = append(grpInfo, info)
            gaccinfo[accountId] = grpInfo
        } else {
            grp = append(grp, info)
            gaccinfo[accountId] = grp
        }
    }
    printMap();
}

func ParseOpt(fileName *string) {
    flag.StringVar(fileName, "i", " ", "inputfile")
    flag.Parse()
    fmt.Println("fileName: ", *fileName);
}


func main() {
    var fileName string;
    ParseOpt(&fileName);
    readFile(fileName);
}
