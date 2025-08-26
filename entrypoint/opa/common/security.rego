package common.security
import rego.v1

# Suspicious path checks
deny[msg] if {
    contains(input.path, "..")
    msg := "path contains .."
}

deny[msg] if {
    contains(input.path, "//")
    msg := "path contains //"
}

# Additional security checks can be added here
deny[msg] if {
    contains(input.path, "../")
    msg := "path contains ../"
}

deny[msg] if {
    regex.match(`\.\./`, input.path)
    msg := "path contains directory traversal pattern"
}