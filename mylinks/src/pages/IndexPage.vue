<template>

  <LinkComponent v-for="entry in entries" :key="entry.id" v-bind="entry" class="q-mb-lg"/>

  <q-dialog v-model="addOrEditEntry" @before-hide="edit = false" persistent>

    <q-card style="width: 50%">
      <q-card-section class="row items-center">
        <div class="text-h6">Add New</div>
        <q-space />
        <q-btn icon="close" flat round dense v-close-popup />
      </q-card-section>

      <q-separator />

      <q-card-section>
        <q-form @submit="submitEntry">
          <q-input label="Title" v-model="entryForm.title" />
          <q-input
            label="Category"
            v-model="entryForm.category"
            style="white-space: pre-wrap"
          />
          <q-input
            label="Description"
            v-model="entryForm.description"
            type="textarea"
            autogrow
          />
          <q-input
            label="URL"
            v-model="entryForm.url"
            type="url"
            style="white-space: pre-wrap"
          />
          <q-btn
            label="Save"
            class="full-width q-mt-lg"
            color="primary"
            type="submit"
          ></q-btn>
        </q-form>
      </q-card-section>
    </q-card>

  </q-dialog>

  <q-dialog v-model="confirmDelete" />

  <q-page-sticky position="bottom-right" :offset="[18, 18]">
    <q-btn fab icon="add" color="primary" @click="openNewDialog" />
  </q-page-sticky>
</template>

<script setup lang="ts">
import {
  addDoc,
  collection,
  doc,
  getDocs,
  updateDoc,
} from 'firebase/firestore';
import { onMounted, ref } from 'vue';
import { db } from 'boot/firebase';
// import { useQuasar } from 'quasar';
import { LinkModel } from 'src/model/LinkModel';
import LinkComponent from 'components/LinkComponent.vue';

const LINKS_PATH = 'links';
defineOptions({
  name: 'IndexPage',
});

onMounted(async () => {
  await getEntries();
});

let entries = ref<LinkModel[]>([]);
// const filter = ref('');
const addOrEditEntry = ref(false);
const confirmDelete = ref(false);
let edit = ref(false);
// const $q = useQuasar();

const entryForm = ref<LinkModel>({
  title: '',
  description: '',
  url: '',
  category: '',
  date: new Date(),
});

async function getEntries() {
  entries.value = [];
  const data = await getDocs(collection(db, LINKS_PATH));

  data.docs.forEach((doc) => {
    let entry: LinkModel = { ...(doc.data() as LinkModel) };
    entry.id = doc.id;
    entries.value.push(entry);
  });
}

function openNewDialog() {
  addOrEditEntry.value = true;
  resetEntryForm();
}

async function submitEntry() {
  if (edit.value) {
    const docRef = doc(db, LINKS_PATH, entryForm.value.id as string);
    await updateDoc(docRef, entryForm.value);
  } else {
    const coll = collection(db, LINKS_PATH);
    await addDoc(coll, entryForm.value);
  }

  await getEntries();
  resetEntryForm();
  addOrEditEntry.value = false;
}

function resetEntryForm() {
  entryForm.value = {
    id: '',
    title: '',
    description: '',
    url: '',
    category: '',
    date: new Date(),
  };
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any

</script>
