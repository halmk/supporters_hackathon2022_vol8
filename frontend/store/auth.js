export const state = () => ({
  userUid: '',
  userName: '',
})
export const getters = {
  getUserUid(state) {
    return state.userUid
  },
  getUserName(state) {
    return state.userName
  },
  isLoggedIn(state) {
    return !!state.userName
  },
}

export const mutations = {
  setUserUid(state, userUid) {
    state.userUid = userUid
  },
  setUserName(state, userName) {
    state.userName = userName
  },
}

export const actions = {
  login({ commit }) {
    console.log('login action')
    const provider = new this.$fireModule.auth.GoogleAuthProvider()
    this.$fire.auth
      .signInWithPopup(provider)
      .then((result) => {
        const user = result.user
        console.log('success : ' + user.uid + ' : ' + user.displayName)
        commit('setUserUid', user.uid)
        commit('setUserName', user.displayName)
        console.log(user)
        localStorage.name = user.displayName
        localStorage.uuid = user.uid
        console.log(localStorage)
      })
      .catch((error) => {
        const errorCode = error.code
        console.log('error : ' + errorCode)
      })
  },
  logout({ commit }) {
    console.log('logout action')
    this.$fire.auth.signOut()
    commit('setUserUid', '')
    commit('setUserName', '')
  },
}
