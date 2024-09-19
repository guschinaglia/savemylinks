<template>
  <q-layout view="lHh Lpr lFf">
    <q-page-container>
      <div style="margin: 20%">
        <q-card>
          <q-form @submit="onSubmit">
            <q-card-section>
              <q-input label="Email" v-model="form.email" dense />
              <q-input
                label="Password"
                v-model="form.password"
                dense
                type="password"
              />
            </q-card-section>
            <q-card-actions>
              <q-btn color="primary" class="full-width" type="submit"
                >Login
              </q-btn>
            </q-card-actions>
          </q-form>
        </q-card>
      </div>
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { reactive } from 'vue';
import { signInWithEmailAndPassword } from 'firebase/auth';
import { firebaseAuth } from 'boot/firebase';
import { useRouter } from 'vue-router';

const router = useRouter();

const form = reactive({
  email: '',
  password: '',
});

async function onSubmit() {
  await signInWithEmailAndPassword(
    firebaseAuth,
    form.email,
    form.password
  ).catch((reason) => {
    alert(reason);
    return;
  });

  await router.push('/main');
}
</script>

<style scoped></style>
