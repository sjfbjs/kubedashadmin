import request from '@/utils/request'

export function getDeployments(namespace) {
  return request({
    url: '/kube/deployments/' + namespace,
    method: 'get'
  })
}

export function getNameSpaces() {
  return request({
    url: '/kube/namespaces',
    method: 'get'
  })
}
