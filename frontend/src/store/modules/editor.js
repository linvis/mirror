import { updateEditorCatalog } from '@/api/editor'

const state = {
  catalog: [{
    id: 0,
    key: '0',
    parent: '0',
    metadata: {
      title: 'root',
      level: 0,
      filetype: 'nb',
      tag: []
    },
    reminder: {
      enable: false,
      count: 0,
      last_time: 0,
      next_time: 0
    },
    children: [{
      id: 1,
      key: '1',
      parent: '0',
      metadata: {
        title: 'file1',
        level: 1,
        filetype: 'md',
        tag: []
      },
      reminder: {
        enable: true,
        count: 0,
        last_time: 0,
        next_time: 0
      }
    }, {
      id: 2,
      key: '2',
      parent: '0',
      metadata: {
        title: 'file2',
        level: 1,
        filetype: 'md',
        tag: []
      },
      reminder: {
        enable: false,
        count: 0,
        last_time: 0,
        next_time: 0
      }
    }
    ]
  }],
  activeFile: {
    id: 1,
    key: '1',
    parent: '0',
    metadata: {
      title: 'file1',
      level: 1,
      filetype: 'md',
      tag: []
    },
    reminder: {
      enable: true,
      count: 0,
      last_time: 0,
      next_time: 0
    }
  },
  contents: []
}

const mutations = {
  CHANGE_CATALOG: (state, catalog) => {
    state.catalog = catalog
  },
  CHANGE_ACTIVEFILE: (state, activeFile) => {
    state.activeFile = activeFile
  },
  CHANGE_FILECONTENT: (state, { key, content }) => {
    state.contents[key] = content
  },
  SUBMIT_CATALOG: state => {
    updateEditorCatalog(state.catalog).then(response => {
    })
  }
}
const actions = {
  changeCatalog({ commit }, data) {
    commit('CHANGE_CATALOG', data)
  },
  changeActiveFile({ commit }, data) {
    commit('CHANGE_ACTIVEFILE', data)
  },
  submitCatalog({ commit }) {
    commit('SUBMIT_CATALOG')
  },
  changeFileContent({ commit }, data) {
    commit('CHANGE_FILECONTENT', data)
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

