# Argv

Argv is a library for [Go](https://golang.org) to split command line string into arguments array. 

# Example
```Go
fmt.Printf("%#v\n", argv.MustParse(`ls -al | tee "list of files.txt"`))
// [][]string{[]string{"ls", "-al"}, []string{"tee", "list of files.txt"}}

cmds := argv.NewCommands(argv.MustParse("ls -al | `openssl` md5", argv.Unbackquote)...)
// [/bin/ls -al /usr/bin/openssl md5]

// run command using standard input, output and error:
cmds.Std().Run()

// using custom stdout:
var buf bytes.Buffer
cmds.Out(&buf).Run()
fmt.Println(strings.TrimSpace(buf.String()))
```

# LICENSE
MIT.
