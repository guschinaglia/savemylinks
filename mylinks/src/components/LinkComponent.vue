<template>
  <q-card>
    <q-card-section>
      <div class="row">
        <q-chip color="primary" text-color="white" dense>{{ props.category }}</q-chip>
<!--        <div class="text-accent">{{ props.date }}</div>-->
        <q-space />
        <div class="text-h6">{{ props.title }}</div>
        <q-space />
        <q-btn
          round
          flat
          icon="delete"
          color="red-4"
          @click="deleteEntry"
        />
      </div>
    </q-card-section>
    <q-card-section>
      <q-input type="textarea" autogrow outlined v-model="props.description" />
    </q-card-section>
    <q-card-section>
      <a :href="props.url">{{ props.url }}</a>
    </q-card-section>
  </q-card>
</template>

<script setup lang="ts">
import { LinkModel } from 'src/model/LinkModel';

const props = defineProps<LinkModel>();

async function deleteEntry() {
  console.log('DELETED');
  // $q.dialog({
  //   title: 'Confirm',
  //   message: 'Please confirm delete of item' + row.item_name,
  //   cancel: true,
  // })
  //   .onOk(async () => {
  //     const docRef = doc(db, 'prices', row.id as string);
  //     await deleteDoc(docRef);
  //     await getEntries();
  //   })
  //   .onCancel(() => {
  //     confirmDelete.value = false;
  //   })
  //   .onDismiss(() => {
  //     confirmDelete.value = false;
  //   });
}
</script>

