import request from '@/utils/request'
export function getDeployments(data) {
    return request({
      url: '/kube/deployments',
      method: 'get',
      data
    })
  }
  