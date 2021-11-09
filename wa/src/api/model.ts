interface DbModel {
  id: number;
  /* UTC time (with offset 0 hour) */
  inserted_at: string;
  /* UTC time (with offset 0 hour) */
  updated_at: string;
}

export interface Post extends DbModel {
  title: string;
  body: string;
  published: boolean;
  /* UTC time (with offset 0 hour) */
  published_at: string | null;
  view_name: string;
  views: number;
  tags: Tag[];
}

export interface Tag extends DbModel {
  name: string;
}

export interface Collection extends DbModel { }

export interface User {
  email: string;
  password: string;
}

export interface Credential {
  access_token: string;
}

export interface Image { }
