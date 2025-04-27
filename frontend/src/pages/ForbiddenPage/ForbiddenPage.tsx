import styles from './ForbiddenPage.module.css';

export default function ForbiddenPage() {
  return (
    <div className={styles.notFound}>
      <h1>403 - Forbidden</h1>
      <p>You do not have permission to access this page.</p>
    </div>
  );
}
