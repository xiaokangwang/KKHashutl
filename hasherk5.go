package main

   import (
       "fmt"
       "io"
       "os"
       "math"
       "golang.org/x/crypto/sha3"
   )

   const filechunk = 8192    // we settle for 8KB

   func main() {



       file, err := os.Open(os.Args[1])

       if err != nil {
          panic(err.Error())
       }

      defer file.Close()

      // calculate the file size
      info, _ := file.Stat()

      filesize := info.Size()

      blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

      hash := sha3.New512()

      for i := uint64(0); i < blocks; i++ {
          blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
          buf := make([] byte, blocksize)

          file.Read(buf)
          io.WriteString(hash, string(buf))   // append into the hash
      }

      fmt.Printf("%s checksum is %X\n",file.Name(), hash.Sum(nil))

 }
