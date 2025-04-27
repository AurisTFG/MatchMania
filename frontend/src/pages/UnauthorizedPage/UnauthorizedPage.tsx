import styles from './UnauthorizedPage.module.css';

export default function UnauthorizedPage() {
  return (
    <div className={styles.notFound}>
      <h1>401 - Unauthorized</h1>
      <p>You do not have permission to access this page.</p>
    </div>
  );
}
