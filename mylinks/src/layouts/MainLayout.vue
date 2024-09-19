<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <!--        <q-btn-->
        <!--          flat-->
        <!--          dense-->
        <!--          round-->
        <!--          icon="menu"-->
        <!--          aria-label="Menu"-->
        <!--          @click="toggleLeftDrawer"-->
        <!--        />-->

        <q-toolbar-title> {{ barTitle }}</q-toolbar-title>

        <div>
          <q-btn @click="doSignOut" color="white" text-color="blue"
            >SignOut</q-btn
          >
        </div>
      </q-toolbar>
    </q-header>

    <!--    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>-->
    <!--      <q-list>-->
    <!--        <q-item-label header> Essential Links </q-item-label>-->
    <!--      </q-list>-->
    <!--    </q-drawer>-->

    <q-page-container>
      <div class="q-pa-lg">
        <router-view />
      </div>
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { signOut } from 'firebase/auth';
import { useRouter } from 'vue-router';
import { firebaseAuth } from 'boot/firebase';

const router = useRouter();
const barTitle = '< / >';

defineOptions({
  name: 'MainLayout',
});

// const leftDrawerOpen = ref(false);

// function toggleLeftDrawer() {
//   leftDrawerOpen.value = !leftDrawerOpen.value;
// }

function doSignOut() {
  signOut(firebaseAuth);
  router.push({ name: 'login' });
}
</script>
