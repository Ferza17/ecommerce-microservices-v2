package common.audit
import rego.v1

# Audit log as a set of messages
log[msg] if {
    msg := sprintf("User %s (%s) attempted %s %s", [
        input.user_id,
        input.user_role,
        input.method,
        input.path,
    ])
}