use axum::{routing::post, Extension, Json, Router};
use serde::{Deserialize, Serialize};

use crate::{
    cfg::Configuration,
    core::security::{generate_app_token, AuthenticatedOrganizationWithApplication},
    error::{HttpError, Result},
    v1::utils::api_not_implemented,
};

#[derive(Deserialize, Serialize)]
pub struct DashboardAccessOut {
    pub url: String,
    pub token: String,
}

const SVIX_URL: &str = "https://app.svix.com";

async fn dashboard_access(
    Extension(cfg): Extension<Configuration>,
    AuthenticatedOrganizationWithApplication { permissions, app }: AuthenticatedOrganizationWithApplication,
) -> Result<Json<DashboardAccessOut>> {
    let token = generate_app_token(&cfg.jwt_secret, permissions.org_id, app.id.clone())?;

    let login_key = serde_json::to_vec(&serde_json::json!({
        "appId": app.id,
        "token": token,
        // Region is unused
        "region": "eu"
    }))
    .map_err(|_| HttpError::internal_server_errer(None, None))?;

    let login_key = base64::encode(&login_key);

    // Included for API compatibility, but this URL will not be useful
    let url = format!("{}/login#key={}", SVIX_URL, login_key);

    Ok(Json(DashboardAccessOut { url, token }))
}

pub fn router() -> Router {
    Router::new()
        .route("/auth/dashboard-access/:app_id/", post(dashboard_access))
        .route("/auth/logout/", post(api_not_implemented))
}