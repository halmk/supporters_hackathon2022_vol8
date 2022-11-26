export default function ({ store }) {
  console.log('ミドルウェア')
  store.commit('auth/setUserUid', localStorage.uuid)
  store.commit('auth/setUserName', localStorage.name)
  // store.dispatch('userUid',localStorage.uuid)
  // store.dispatch('userName',localStorage.name)
}
