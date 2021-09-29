interface DbModel {
  id: number;
  insertedAt: string;
  updatedAt: string;
}

export interface Post extends DbModel {
  title: string;
  body: string;
  published: boolean;
  publishedAt: string | null;
  viewName: string;
  views: number;
  tags: string[];
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
  accessToken: string;
}

export interface Image { }
