export const state = () => ({
  userUid: '',
  userName: '',
  userPhoto: '',
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
  getUserPhoto(state) {
    return state.userPhoto
  },
}

export const mutations = {
  setUserUid(state, userUid) {
    state.userUid = userUid
  },
  setUserName(state, userName) {
    state.userName = userName
  },
  setUserPhoto(state, userPhoto) {
    state.userPhoto = userPhoto
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
        commit('setUserPhoto', user.photoURL)
        console.log(user)
        localStorage.name = user.displayName
        localStorage.uuid = user.uid
        localStorage.photo = user.photoURL
        console.log(localStorage)
        const db = this.$fire.firestore
        db.collection('users').doc(user.uid).set({
          name: user.displayName,
        })
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
    commit('setUserPhoto', '')
  },
}
