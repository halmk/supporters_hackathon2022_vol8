export default function ({ store, route, redirect }) {
  console.log('ミドルウェア')
  store.commit('auth/setUserUid', localStorage.uuid)
  store.commit('auth/setUserName', localStorage.name)

  if (route.name === 'projects-project') {
    console.log(route.name)
    const uid = store.getters['auth/getUserUid']
    if (uid === '' || uid === undefined) {
      // router.push('/projects')
      return redirect('/projects')
    }
  }
}
