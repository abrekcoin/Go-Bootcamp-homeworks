package main
/* 

PS C:\Users\faiks\OneDrive\Masaüstü\learngo-master> go doc runtime NumCPU      
package runtime // import "runtime"

func NumCPU() int
    NumCPU returns the number of logical CPUs usable by the current process.

    The set of available CPUs is checked by querying the operating system at
    process startup. Changes to operating system CPU allocation after process
    startup are not reflected.

PS C:\Users\faiks\OneDrive\Masaüstü\learngo-master> 

PS C:\Users\faiks\OneDrive\Masaüstü\learngo-master> go doc -src runtime NumCPU
package runtime // import "runtime"

// NumCPU returns the number of logical CPUs usable by the current process.
//
// The set of available CPUs is checked by querying the operating system
// at process startup. Changes to operating system CPU allocation after
// process startup are not reflected.
func NumCPU() int {
        return int(ncpu)
} */