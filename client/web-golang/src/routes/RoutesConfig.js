import Login from "../containers/login/Login";
import DocumentManagement from "../containers/document_management/DocumentManagement";


const LOGIN = { container: Login, path: "/login", isPrivate: false };
const DOCUMENT_MANAGEMENT = { container: DocumentManagement, path: "/pnl", isPrivate: false };


export default [LOGIN, DOCUMENT_MANAGEMENT];