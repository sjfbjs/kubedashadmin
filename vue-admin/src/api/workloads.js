import request from '@/utils/request'
export function getDeployments(params) {
    return request({
      url: '/kube/deployments',
      method: 'get',
      params
    })
  }
  