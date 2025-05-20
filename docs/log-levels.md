# Logging system log levels

|Log Level| Description |
|--|--|
| Trace | Super detailed debugging info (e.g., every function call, variable change) |
| Debug | General debug info, useful during development |
| Info  | Normal operational messages (startup, shutdown, requests handled) |
| Warn  | Something unexpected happened, but the app can recover |
| Error | Something failed, but the app is still running |
| Fatal | Logs the error and exits the application with os.Exit(1) |
| Panic | Logs the error and panics, causing a stack trace |

## Trace use cases
- Entering/exiting functions
- internal loops 
- step-by-step operations

## Debug use cases
- detailed info	
- Validation failures
- variable values
- decisions in code flow

## Info use cases
- important events
- User logins
- successful transactions
- system start/shutdown

## Warn use cases
- Retryable errors
- slow response times
- deprecated APIs

## Error use cases
- Failed database writes
- panics
- unauthorized access attempts

## Fatal use cases
- Crash
- unrecoverable state