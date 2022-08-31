# -*- mode: Python -*-
load('ext://restart_process', 'docker_build_with_restart')
local_resource(
    'auth-compile',
    cmd='cd auth; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/auth main.go',
    deps=['auth/main.go', 'auth/security', 'auth/user', 'pkg'],
)

docker_build_with_restart(
    'auth-image',
    './auth',
    dockerfile='auth/Dockerfile',
    entrypoint=['/auth'],
    live_update=[
        sync('./auth/bin/auth', '/auth'),
    ],
)

k8s_yaml('auth/kubernetes.yaml')
k8s_resource('ms-auth', port_forwards=8081,
             resource_deps=['auth-compile'])


local_resource(
    'feedbacks-compile',
    cmd='cd feedbacks; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/feedbacks main.go',
    deps=['feedbacks/main.go', 'feedbacks/feedback', 'pkg'],
)

docker_build(
    'feedbacks-image',
    './feedbacks',
    dockerfile='feedbacks/Dockerfile',
)

local_resource(
    'votes-compile',
    cmd='cd votes; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/votes main.go',
    deps=['votes/main.go', 'votes/vote', 'pkg'],
)

docker_build(
    'votes-image',
    './votes',
    dockerfile='votes/Dockerfile',
)
