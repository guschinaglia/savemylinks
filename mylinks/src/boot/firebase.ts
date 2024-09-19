import { initializeApp } from 'firebase/app';
import { getAuth } from 'firebase/auth';
import { getFirestore } from 'firebase/firestore';

const firebaseConfig = {
  apiKey: 'AIzaSyBiKOfCyUqrDhsmUZsAGM-ong3UuH1pHKs',
  authDomain: 'my-links-16511.firebaseapp.com',
  projectId: 'my-links-16511',
  storageBucket: 'my-links-16511.appspot.com',
  messagingSenderId: '1061400497684',
  appId: '1:1061400497684:web:6fe51002731e8d8e8ca101',
};

const firebaseApp = initializeApp(firebaseConfig);
const firebaseAuth = getAuth(firebaseApp);
const db = getFirestore(firebaseApp);

export { firebaseApp, firebaseAuth, db };
