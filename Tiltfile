# -*- mode: Python -*-
# compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/poc-api ./cmd/api/main.go'
# local_resource(
#     'example-go-compile',
#     compile_cmd,
#     deps=['./cmd/main.go']
# )

local_resource('auth', cmd='cd auth; go build -o bin/auth main.go',
               serve_cmd='auth/bin/auth', deps=['auth/main.go', 'auth/security', 'auth/user', 'pkg'])

local_resource('feedbacks', cmd='cd feedbacks; go build -o bin/feedbacks main.go',
               serve_cmd='feedbacks/bin/feedbacks', deps=['feedbacks/main.go', 'feedbacks/feedback', 'pkg'])


local_resource('votes', cmd='cd votes; go build -o bin/votes main.go',
               serve_cmd='votes/bin/votes', deps=['votes/main.go', 'votes/vote', 'pkg'])
