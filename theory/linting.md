# golangci-lint
https://golangci-lint.run/welcome/install/#local-installation

https://golangci-lint.run/welcome/quick-start/

 $ sudo snap install golangci-lint --classic
 golangci-lint 1.64.5 from GolangCI-Lint (golangci✓) installed

 $ golangci-lint --version
 golangci-lint has version 1.64.5 built with go1.24.0 from 0a603e49 on 2025-02-13T21:19:55Z


## enabled or disabled linters?
  $ golangci-lint help linters

## run the linter

- Run linting across all files within the root directory.

learning-golang$ golangci-lint run
mypackages/feedback/drinking.go:6:31: illegal rune literal (typecheck)
        } else if age >= 18 && age < '80' {
                                     ^
main.go:15:17: undefined: greetings (typecheck)
    fmt.Println(greetings.GreetingScroll(name))
                ^
main.go:16:14: undefined: feeback (typecheck)
        fmt.Println(feeback.DrinkingAge(age))
                    ^
main.go:17:34: undefined: ae (typecheck)
        fmt.Println(feedback.DrivingAge(ae))
                                        ^

- Run in specific folder 

learning-golang$ golangci-lint run mypackages/feedback/
mypackages/feedback/drinking.go:6:31: illegal rune literal (typecheck)
        } else if age >= 18 && age < '80' {
