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

docker_build_with_restart(
    'feedbacks-image',
    './feedbacks',
    dockerfile='feedbacks/Dockerfile',
    entrypoint=['/feedbacks'],
    live_update=[
        sync('./feedbacks/bin/feedbacks', '/feedbacks'),
    ],
)

k8s_yaml('feedbacks/kubernetes.yaml')
k8s_resource('ms-feedbacks', port_forwards=8082,
             resource_deps=['feedbacks-compile'])


local_resource(
    'votes-compile',
    cmd='cd votes; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/votes main.go',
    deps=['votes/main.go', 'votes/vote', 'pkg'],
)

docker_build_with_restart(
    'votes-image',
    './votes',
    dockerfile='votes/Dockerfile',
    entrypoint=['/votes'],
    live_update=[
        sync('./votes/bin/votes', '/votes'),
    ],
)

k8s_yaml('votes/kubernetes.yaml')
k8s_resource('ms-votes', port_forwards=8083,
             resource_deps=['votes-compile'])
