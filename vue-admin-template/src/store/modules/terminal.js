// const state = {
//   sshParams: null
// }
// const mutations = {
//   setSSHParams(state, params) {
//     state.sshParams = params
//   }
// }

// export default {
//   namespaced: true, // 启用命名空间
//   state,
//   mutations
// }

export default {
  namespaced: true,
  state: () => ({
    sshParams: null
  }),
  mutations: {
    setSSHParams(state, params) {
      state.sshParams = params
    }
  },
  getters: {
    formattedParams: state => {
      return state.sshParams
        ? `${state.sshParams.username}@${state.sshParams.host}`
        : '未设置'
    }
  }
}
