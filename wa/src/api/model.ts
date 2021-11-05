interface DbModel {
  id: number;
  inserted_at: string;
  updated_at: string;
}

export interface Post extends DbModel {
  title: string;
  body: string;
  published: boolean;
  published_at: string | null;
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
  access_token: string;
}

export interface Image { }
