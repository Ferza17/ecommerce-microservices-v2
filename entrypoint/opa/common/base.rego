package common.base
default allow = false

# --- audit log as a set of messages
audit_log[msg] if {
    msg := sprintf("User %s (%s) attempted %s %s", [
        input.user_id,
        input.user_role,
        input.method,
        input.path,
    ])
}


# --- suspicious path checks
deny[msg] if {
    contains(input.path, "..")
    msg := "path contains .."
}

deny[msg] if {
    contains(input.path, "//")
    msg := "path contains //"
}